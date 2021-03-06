#!/bin/zsh

# Script used to build go packages for multipule OS/Arch's 

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_name=$package

#the full list of the platforms: https://golang.org/doc/install/source#environment
# Comment in or out as needed
platforms=(
	"darwin/amd64"

    # "aix/ppc64"
	# "android/386"
	# "android/amd64"
	# "android/arm"
	# "android/arm64"
	# "darwin/arm64"
	# "dragonfly/amd64"
	# "freebsd/386"
	# "freebsd/amd64"
	# "freebsd/arm"
	# "illumos/amd64"
	# "js/wasm"
	# "linux/386"
	# "linux/amd64"
	# "linux/arm"
	# "linux/arm64"
	# "linux/ppc64"
	# "linux/ppc64le"
	# "linux/mips"
	# "linux/mipsle"
	# "linux/mips64"
	# "linux/mips64le"
	# "linux/s390x"
	# "netbsd/386"
	# "netbsd/amd64"
	# "netbsd/arm"
	# "openbsd/386"
	# "openbsd/amd64"
	# "openbsd/arm"
	# "openbsd/arm64"
	# "plan9/386"
	# "plan9/amd64"
	# "plan9/arm"
	# "solaris/amd64"
	# "windows/386"
	# "windows/amd64 "
)

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
