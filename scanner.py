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
""")

print("selected option: ", response)

if response == '1':
    print("nmap version: ", scanner.nmap_version())
    scanner.scan(ip_address, '1-1024', -'v -s5')
    print(scanner.scaninfo())