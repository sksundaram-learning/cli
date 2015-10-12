%{define} erisbuilddir /var/tmp/eris-rpmbuild.tmp

Summary: Eris is a platform for building, testing, maintaining, and operating distributed applications with a blockchain backend. Eris makes it easy and simple to wrangle the dragons of smart contract blockchains.
Name: eris-tools
Version: 0.10.4
Release: 1
License: GPL-3
Group: Applications/Productivity
URL: https://docs.erisindustries.com
BuildRoot: buildroot-%{name}-%{version}-%{release}.%{_arch}
BuildRequires: golang
Requires: docker

%description 
eris tools to manipulate blockchains smartcontracts

%prep
rm -fr %{erisbuilddir}
mkdir %{erisbuilddir}
git clone https://github.com/eris-ltd/eris-cli %{erisbuilddir}

%build
pushd %{erisbuilddir}/cmd/eris
GOPATH=%{erisbuilddir} go get
GOPATH=%{erisbuilddir} go build
popd

%install
mkdir -p ${RPM_BUILD_ROOT}/%{_bindir}
install %{erisbuilddir}/cmd/eris/eris ${RPM_BUILD_ROOT}/%{_bindir}

%files
%defattr(-, root, root, 0755)
%{_bindir}/*

