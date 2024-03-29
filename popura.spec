Name:           popura
Version:        0.4.3+popura1
Release:        1%{?dist}
Summary:        Popura ポプラ: alternative Yggdrasil network client

License:        GPL-3.0-only
Group:          Productivity/Networking/Other
URL:            https://popura-network.github.io
Source:         https://codeload.github.com/popura-network/Popura/tar.gz/v%{version}

%{?systemd_requires}
BuildRequires:  systemd go >= 1.16 git
Requires(pre):  shadow
Conflicts:      popura-develop yggdrasil yggdrasil-develop

%description
Popura uses the same Yggdrasil core API internally, but adds some useful and
experimental features which the original client lacks.

%define debug_package %{nil}

%prep
%setup -qn Popura-%(echo %{version} | sed -e 's/+/-/g')

%build
export PKGNAME="%{name}"
export PKGVER="%{version}"
export GOPROXY="https://proxy.golang.org,direct"
./build -t -p -l "-linkmode=external"

%install
rm -rf %{buildroot}
install -m 0755 -D yggdrasil %{buildroot}/%{_bindir}/yggdrasil
install -m 0755 -D yggdrasilctl %{buildroot}/%{_bindir}/yggdrasilctl
install -m 0644 -D contrib/systemd/yggdrasil.service %{buildroot}%{_prefix}/lib/systemd/system/yggdrasil.service
install -m 0644 -D contrib/systemd/yggdrasil-default-config.service %{buildroot}%{_prefix}/lib/systemd/system/yggdrasil-default-config.service
mkdir -p %{buildroot}%{_sbindir}
ln -sf service %{buildroot}%{_sbindir}/rcyggdrasil

%files
%{_bindir}/yggdrasil
%{_bindir}/yggdrasilctl
%{_prefix}/lib/systemd/system/yggdrasil.service
%{_prefix}/lib/systemd/system/yggdrasil-default-config.service
%{_sbindir}/rcyggdrasil

%pre
getent group yggdrasil >/dev/null || groupadd -r yggdrasil
%service_add_pre yggdrasil.service

%post
%service_add_post yggdrasil.service

%preun
%service_del_preun yggdrasil.service

%postun
%service_del_postun yggdrasil.service

%changelog
