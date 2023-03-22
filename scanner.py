#nmap automation
import nmap

scanner = nmap.PortScanner()

print("welcome, ")
print("<------------------------------------------->")

ip_address = input("enter IP address: ")
print("IP entered is: ", ip_address)
type(ip_address)

response = input("""\n Enter the type of scan
            1)SYNC ACK Scan
            2)UDP San
            3)Comprehensive Scan
\n""")

print("selected option: ", response)

if response == '1':
    print("nmap version: ", scanner.nmap_version())
    scanner.scan(ip_address, '1-1024', '-v -sS')
    print(scanner.scaninfo())
    print("ip status: ", scanner[ip_address].state())
    print(scanner[ip_address].all_protocols())
    print("open ports: ", scanner[ip_address]['tcp'].keys())
