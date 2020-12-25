package autopeering
var PublicPeers = []string{
	"tcp://104.248.15.125:31337",
	"tcp://108.175.10.127:61216",
	"tcp://139.162.119.37:44478",
	"tcp://140.238.168.104:17117",
	"tcp://176.101.222.31:27363",
	"tcp://176.215.237.83:2755",
	"tcp://176.223.130.120:22632",
	"tcp://188.226.125.64:54321",
	"tcp://194.177.21.156:5066",
	"tcp://195.123.245.146:7743",
	"tcp://198.58.100.240:44478",
	"tcp://[2001:19f0:5401:272a:5400:2ff:fe21:7a26]:8082",
	"tcp://[2001:19f0:7402:431:5400:2ff:fe21:7912]:8082",
	"tcp://[2001:1af8:4700:a119:7::1]:35239",
	"tcp://[2001:41d0:2:c44a:51:255:223:60]:26409",
	"tcp://[2001:41d0:401:3000::4227]:62506",
	"tcp://[2001:41d0:601:1100::cf2]:37145",
	"tcp://[2001:8d8:1800:8224::1]:61216",
	"tcp://[2001:bc8:1820:192f::1]:62486",
	"tcp://212.129.52.193:39565",
	"tcp://217.195.164.4:10000",
	"tcp://[2400:8902::f03c:91ff:fe1f:c32a]:44478",
	"tcp://[2600:3c00::f03c:91ff:feae:3efa]:44478",
	"tcp://[2604:a880:400:d0::16e5:7001]:19103",
	"tcp://[2604:a880:800:c1::2c2:a001]:31337",
	"tcp://[2607:f1c0:1801:d4::1]:61216",
	"tcp://[2804:49fc::ffff:ffff:5b5:e8be]:58301",
	"tcp://[2a00:b700:2::6:69]:1333",
	"tcp://[2a02:17d0:1b4:bddd::7]:54321",
	"tcp://[2a03:3b40:fe:ab::1]:46370",
	"tcp://[2a03:6f00:5:1::59df:79d3]:5353",
	"tcp://[2a05:3580:d900:1b13:e2d5:5eff:fed8:8b86]:8777",
	"tcp://[2a05:9403::8b]:7743",
	"tcp://37.205.14.171:46370",
	"tcp://45.11.19.26:5001",
	"tcp://45.231.133.188:58301",
	"tcp://45.76.137.140:8082",
	"tcp://45.77.107.150:34660",
	"tcp://46.105.92.61:62506",
	"tcp://46.151.26.194:60575",
	"tcp://50.236.201.218:56088",
	"tcp://51.15.118.10:62486",
	"tcp://51.15.204.214:12345",
	"tcp://51.255.223.60:26409",
	"tcp://51.75.44.73:50001",
	"tcp://54.37.137.221:37145",
	"tcp://64.112.176.176:1617",
	"tcp://64.112.177.94:1617",
	"tcp://64.112.180.77:1617",
	"tcp://64.112.182.119:1617",
	"tcp://64.112.182.182:1617",
	"tcp://67.205.187.55:19103",
	"tcp://78.27.153.163:33165",
	"tcp://82.165.69.111:61216",
	"tcp://85.17.15.221:35239",
	"tcp://88.201.129.205:8777",
	"tcp://89.223.121.211:5353",
	"tcp://91.206.93.65:1333",
	"tcp://95.165.99.143:61933",
	"tcp://ams1.y.sota.sh:8080",
	"tcp://curiosity.tdjs.tech:30003",
	"tcp://edge.v4.hel1.devices.y.samip.fi:65444",
	"tcp://edge.v4.tku1.devices.y.samip.fi:65444",
	"tcp://edge.v6.hel1.devices.y.samip.fi:65444",
	"tcp://edge.v6.tku1.devices.y.samip.fi:65444",
	"tcp://glimmy.092918.xyz:50001",
	"tcp://hindsight.krvtz.net:8082",
	"tcp://lancis.iscute.moe:49273",
	"tcp://lan.tdem.in:50001",
	"tcp://n2o.ddns.net:22632",
	"tcp://nessie.krvtz.net:8082",
	"tcp://vps.tomasgl.ru:61933",
	"tcp://yggdrasil.frank2.net:1337",
	"tcp://ygg.loskiq.com:17313",
	"tcp://yggnode.ddns.net:18226",
	"tcp://yggnode.ddns.net:18228",
	"tcp://y.zbin.eu:7743",
	"tls://104.248.15.125:32337",
	"tls://140.238.168.104:17121",
	"tls://176.101.222.31:27364",
	"tls://176.215.237.83:2756",
	"tls://194.177.21.156:5068",
	"tls://[2001:41d0:2:c44a:51:255:223:60]:10259",
	"tls://[2001:41d0:401:3000::4227]:46394",
	"tls://[2001:41d0:601:1100::cf2]:14987",
	"tls://212.129.52.193:39575",
	"tls://217.195.164.4:10531",
	"tls://[2604:a880:400:d0::16e5:7001]:19102",
	"tls://[2604:a880:800:c1::2c2:a001]:32337",
	"tls://[2a00:b700:2::6:69]:1444",
	"tls://[2a01:d0:ffff:4353::2]:6010",
	"tls://[2a02:1802:5e:0:18d2:e2ff:fe44:17d2]:9944",
	"tls://[2a02:1802:5e:0:4b7:c2ff:fe29:ba37]:9944",
	"tls://[2a02:1802:5e:0:7cc7:cdff:fe28:128a]:9944",
	"tls://[2a02:1802:5e:0:8854:12ff:fe35:65b3]:9944",
	"tls://[2a02:1802:5e:0:900c:33ff:fea5:7aec]:9944",
	"tls://[2a02:1802:5e:0:b841:d3ff:feb2:9a81]:9944",
	"tls://[2a02:1802:5e:0:b84d:bdff:fe90:1a82]:9944",
	"tls://[2a02:1802:5e:0:bc20:44ff:fe93:6353]:9944",
	"tls://[2a05:3580:d900:1b13:e2d5:5eff:fed8:8b86]:8778",
	"tls://45.11.19.26:5002",
	"tls://45.147.198.155:6010",
	"tls://46.105.92.61:46394",
	"tls://46.151.26.194:8443",
	"tls://51.255.223.60:10259",
	"tls://54.37.137.221:14987",
	"tls://67.205.187.55:19102",
	"tls://78.27.153.163:33166",
	"tls://88.201.129.205:8778",
	"tls://91.206.93.65:1444",
	"tls://edge.v4.hel1.devices.y.samip.fi:65445",
	"tls://edge.v4.tku1.devices.y.samip.fi:65445",
	"tls://edge.v6.hel1.devices.y.samip.fi:65445",
	"tls://edge.v6.tku1.devices.y.samip.fi:65445",
	"tls://glimmy.092918.xyz:58008",
	"tls://lancis.iscute.moe:49274",
	"tls://lan.tdem.in:50002",
	"tls://vps.tomasgl.ru:61944",
	"tls://ygg.loskiq.com:17314",
	"tls://yggnode.ddns.net:18227",
	"tls://yggnode.ddns.net:18229"}