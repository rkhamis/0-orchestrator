copypkg() {
   pkgname=$1
   target=$2
   for file in $(dpkg -L $pkgname); do
      if [ -d "$file" ]; then
         continue
      fi
      dirname=$(dirname $file)
      if [[ "$dirname" == /usr/share/man* ]]; then
         continue
      fi
      if [[ "$dirname" == /usr/share/doc* ]]; then
         continue
      fi
      targetdir="${target}${dirname}"
      mkdir -p "$targetdir"
      cp $file "$targetdir"
      if file "$file" | grep dynamic; then
         lddcopy "$file" "$target"
      fi
   done
}

ensure_lddcopy() {
   if ! which lddcopy; then
      pushd /tmp
      git clone --depth=1 https://github.com/maxux/lddcopy.git
      cp /tmp/lddcopy/lddcopy.sh /usr/local/bin/lddcopy
      chmod +x /usr/local/bin/lddcopy
      popd
   fi
}

ensure_go() {
   if ! which go; then
      curl https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz > /tmp/go1.8.linux-amd64.tar.gz
      tar -C /usr/local -xzf /tmp/go1.8.linux-amd64.tar.gz
      export PATH=$PATH:/usr/local/go/bin
   fi
   mkdir -p /gopath
   export GOPATH=/gopath

}
