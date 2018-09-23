# AWS note
 
- SIT, system integration test
- UAT, user acceptance test
- Prod, prod
	- AWS environment and provides best practice recommendations in these five categories: 
- Cost Optimization, 
- Performance, 
- Fault Tolerance, 
- Security
- Service Limits.

### Amazon Elastic Compute Cloud, EC2
##### type 
    + reserved instance, 預定機器
        * Discommission: 
            1. Snapshot EBS and terminate
            2. Sell instance in reserved instance marketplace
    + spot instance, 競價機器
    + on demand instance, 按需求開
    + dedicated instance, 專屬實體機器
    Billing commences when Amazon EC2 initiates the boot sequence of an AM instance. Billing ends when the instance shuts down, which could occur through a web services command, by running “shutdown -h”, or through instance failure
##### Failover
    + active-active: two servers. when one is down, other take over
        * pros: sharing load, switch fast, lower RTO, frequent RPO
        * cons: some harmful request will influence the secondary server
    + active-passive: two nodes but only one server. when server is down, restart it on other node.
    + pros: Won't require extra machine, 
    + cons: Need few period of time to boot up or restart. Higher RTO
##### RPO RTO
    Recovery Point Objective: 備份週期 高 -> high cost, 資料完整, lower RTO
        you may determine that 5% of the data must be available within 12 hours, 50% of the data must be available after a complete loss of the database within 2 days, and the remainder of the data be available within 5 days. In this case you have two RTOs. Your total RTO is 7.5 days.
    Recovery Time Objective: 恢復時間 短 -> less impact of service, 
##### Capacity
    + number of instance, reserved limit: 20, on-demand limit is depended on type of instance
    + Key pairs: 5000
    + Placement groups: 500
    + Network interfaces per instance: depend on type of instance
    + secruity group per VPC: 500
    + inbound or outbound rules per security group: 50
    EC2-Classic
    + Elastic IP addresses for EC2-Classic: 5
    + Security groups for EC2-Classic per instance: 500
    + Rules per security group for EC2-Classic: 100
    + Can't change security gruop but addd or remove rules, you can‟t change the outbound rules for EC2-Classic. 
    + In Amazon EC2-Classic every instance is given two IP Addresses: a private IF address and a public IP address
##### types of storage
    + instance storage = EBS-backed
    + EBS

    /dev/sda1, reserved root device name for paravirtual and HVM
    /dev/xvda, reserved root device name for HVM
##### placement group
    Placement groups are recommended for applications that benefit from low network latency, high network throughput, or both. To provide the lowest latency, and the highest packet-per-second network performance for your placement group, choose an instance type that supports enhanced networking
    + Max number of instance is as same as max number of ec2 instance per region
    + EC2 with higher communication capacity. They should be in the same AZ.
    + 10 GPS network. 
    + You can‟t move an existing instance into a placement group. 
    + You can create an AM from your existing instance, and then launch a new instance from the AMI into a placement group.
    + A placement group can span peered VPCs; however, you will not get full-bisection bandwidth between instances in peered VPCs
##### Network Performance
    network performance
        
        10GB => 
            c4.8xlarge, 
            c5.9xlarge, 
            d2.8xlarge, 
            g3.8xlarge, 
            h1.8xlarge, 
            3.8xlarge, 
            m4.10xlarge, 
            m5.12xlarge, 
            p2.8xlarge, 
            p3.8xlarge, 
            r4.8xlarge, 
            x1.16xlarge, 
            x1e.16xlarge
        20GB => 
            g3.16xlarge, 
            c5.18xlarge, 
            h1.16xlarge, 
            i3.16xlarge, 
            m4.16xlarge, 
            m5.24xlarge, 
            p2.16xlarge, 
            p3.16xlarge, 
            r4.16xlarge, 
            x1.32xlarge, 
            x1e.32xlarge
##### Elatic Network Interface, ENI
    An EC2 on two different subnets, ENI should be in the same AZ. 
    It won't double the bandwidth.
##### Elastic IP
    [page](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html#VPC_EIP_EC2_Differences)
    If your account supports EC2-Classic, there's one pool of Elastic IP addresses for use with the EC2-Classic platform and another for use with the EC2-VPC platform. You can't associate an Elastic IP address that you allocated for use with a VPC with an instance in EC2-Classic, and vice- versa. 
    When you associate an Elastic IP address with an instance in EC2-Classic, a default VPC, or an instance in a nondefault VPC in which you assigned a public IPv4 to the eth0 network interface during launch, the instance's current public IPv4 address is released back into the public IP address pool. If you disassociate an Elastic IP address from the instance, the instance is automatically assigned a new public IPv4 address within a few minutes.
    
    However, if you have attached a second network interface to the instance, the instance is not automatically assigned a new public IP address; you‟ll have to associate an EIP with it manual y. The EIP remains associated with the instance.
##### AMI
    AMI is stored at S3
    If the user has deregistered the AMI of an EC2 instance and is trying to launch a similar instance with the option “Launch more like this, AWS will throw an error saying that the AMI is deregistered or not available.
##### TMP?
    
    You can seamlessly join an EC2 instance to your directory domain when the instance is launched using the Amazon EC2 Simple Systems Manager. If you need to manual join an EC2 instance to your domain, you must launch the instance in the proper region and security group or subnet, then join the instance to the domain. To be able to connect remotely to these instances, you must have IP connect Myth to the instances from the network you are connecting from. In most cases, this requires that an Internet gateway be attached to your VPC and that the instance has a public IP address.
    EC2
        procure instances as reserved instances beforehand
    ec2 1 hr everyday = on-demand , spot?

    Security groups are state full, Return traffic is automatically allowed, regardless of any rules
    Network ACLs are stateless, Return traffic must be explicitly allowed by rules
    
    AMI, Elastic IP, AZ can help on HA

### Elastic Container Service, ECS
    Docker is only suppport on ec2 container service
##### Task definition 
    configure file to launch container in ecs

### Amazon Elastic MapReduce, EMR
    Amz EMR state, 
        STARTING — The cluster provisions, starts, and configures EC2 instances. 
        BOOTSTRAPPING — Bootstrap actions are being executed on the cluster. 
        RUNNING — A step for the cluster is currently being run. 
        WAITING — The cluster is currently active, but has no steps to run. 
        TERMINATING - The cluster is in the process of shutting down. 
        TERMINATED - The cluster was shut down without error. 
        TERMINATED_WITH_ERRORS - The cluster was shut down with errors.
    EMR encryption: 
    + SSE-s3 
    + SSE-kms 
    + CES-kms 
    + CES-custom
    5.7~ custome AMI for encrypting EBS root device
    if no security configuration, configure in s3 manually
    4.1~  transparent encryption in HDFS, http://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-hdfs-config.html#emr-encryption-tdehdfs


# Storage
### Storage Gateway
    + gateway-cached volumes, cache in on-premise server
    + gateway-stored volumes , store in on-premise server

### Amazon Glacier, Virtual Tape Shelf
    Notification configuration
    Job
    Archive
##### Glacier Vault
    Up to 1,000 vaults per region. Update around once a day
    A vault is a way to group archives together in Amazon Glacier. You organize your data in Amazon Glacier using vaults. Each archive is stored in a vault of your choice. 
    You may control access to your data by setting vault-level access policies using the AWS Identity and Access Management (IAM) service. You can also attach notification policies to your vaults. 
    These enable you or your application to be notified when data that you have requested for retrieval is ready for download. Click here to learn more about setting up notifications using the Amazon Simple Notification Service (Amazon SNS).

### Amazon Red Shift
    AWS recommends Amazon Red shift for customers who have a combination of needs, such as: 
        - High performance at scale as data
        - Query complexity grows Desire to prevent reporting and analytic processing from interfering with the performance of OLTP workloads Large volumes of structured data to persist and 
        - Query using standard SQL and existing BI tools Desire to the administrative burden of running one‟s own data warehouse and dealing with setup, durability, monitoring, scaling and patching.
    Amazon Red shift achieves efficient storage and optimum query performance through a combination of 
        - massively parallel processing
        - columnar data storage
        - very efficient, targeted data compression encoding schemes. 

### Simeple Storage Service, S3
    S3 doesn't provide search functionality. In this case it is recommended to have an own DB system which manages the S3 metadata and key mapping.

    s3 charge less where we cost less
##### Type
    + Standard
    + Standard-Infrequent Access, Standard-IA
    + One AZ- IA
    Standard and Standard IA replicate data across a minimum of three AZs 
##### Capacity
    + The total volume of data and number of objects you can store are unlimited.
    + Individual Amazon S3 objects can range in size from a minimum of 0 bytes to a maximum of 5 terabytes.  
    + The largest object that can be uploaded in a single PUT is 5 gigabytes. 
    + For objects larger than 100 megabytes, customers should consider using the Multipart Upload capability.
    + Using Bit Torrent protocol, object should be less than 5GB
    + 100 buckets per account
##### Select Query
    
    only select particular portion of Object without return whole object
##### Read after Write
    
    A newly created or updated object will be visible immediately.
#####  Origin Access Identity (OAI)
    A special CloudFront user (OAI) associate your origin in order to protect the origin.
##### Access Control
    AWS IAM policies, 
        You can grant IAM users fine-grained control to your Amazon S3 bucket or objects.
    Access Control List (ACL), 
        You can use ACIs to selectively add (grant) certain permissions on individual objects.
    Resource-Based Policy and User policy
    Bucket Policy
        Amazon S3 bucket policies can be used to add or deny permissions across some or all of the objects within a single bucket. Default policy is private. Only resource owner and AWS account which created the bucket have access right. Wildcard(*) for anonymous user.
    [Query string authentication](https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
        With Query string authentication, you have the ability to share Amazon S3 objects through URIs that are valid for a specified period of time.
    S3 MFA, multi-factor autherization, https://aws.amazon.com/tw/iam/details/mfa/ 
        做動作需要額外授權 類似PingID
##### Crypto
    + Data encryption at rest, encrypts your data as it writes it to disks in its data centers and decrypts it for you when you access it
        
    + Multi-Factor encryption (MFA). Amazon S3 encrypts each object with a unique key. As an additional safeguard, it encrypts the key itself with a master key that it regularly rotates. 
    + Amazon S3 server-side encryption uses one of the strongest block ciphers available, 256-bit Advanced Encryption Standard (AES-256)AWS-KMS, to encrypt your data. (Same as AWS Glacier). SSE-S3 SSE-KMS
    + In client-side encryption, you manage encryption/decryption of your data, the encryption keys, and related tools. CES-KMS, CES-Custom
#####  tmp
    + List your AWS import/Export jobs through S3 REST API

### Amazon Elastic Block Store, EBS
    Resource-based permissions are supported.
    When EBS is as a root device, the instance should be shutdown or stopped.
    The data in an instance store persists only during the lifetime of its associated instance. If an instance reboots (intentionally or unintentionally), data in the instance store persists. However, data on instance store volumes is lost under the following circumstances. 
     AID 5 and RAID 6 are not recommended for Amazon EBS because the parity write operations of these RAID modes consume some of the lops available to your volumes.
##### EBS volumn type
    [IOPS calculation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)
    General Purpose SSD gp2
    1GiB ~ 16TiB, based on 16 KiB/io and 256KiB/io
    Max Throughput: 160MiB/s/vol, 10000iops/vol, 80000iops/ins
    Burst Credit: 3 iops/GiB (max 5.4 million I/O credits = initial credit) 
    Burst Rate: 3 iops/GiB max burst: 3000iops
    Balanced at 1000GiB
    16KiB/iops max (3000x16KiB/1024 = 46.88)MiB/s (Due to max burst rate) 
    256KiB/iops max 160MiB/s(balanced at 214 GiB)
        General Purpose SSD (gp2) volumes have a throughput limit between 128 MiB/s and 160 MiB/s depending on volume size. Volumes greater than 170 GiB up to 214 GiB deliver a maximum throughput of 160 MiB/s if burst credits are available. Volumes above 214 GiB deliver 160 MiB/s irrespective of burst credits.

    Provisioned IOPS SSD io1
    4GiB ~ 16TiB, based on 16 KiB/io and 256KiB/io
    Max Throughput: 500Mib/s/vol, 32000iops/vol, 80000iops/ins
    Burst Credit: 50iops/GiB max(32000iops at 640GiB)
    Burst Rate;
    16KiB/iops max 500MiB/s(32000iops)
    256KiB/iops max 500MiB/s(2000iops)
    Throughput Optimized HDD st1
    500GiB ~ 16TiB, based on 1 MiB/io
    Max Throughput: 500Mib/s/vol, 500iops/vol, 80000iops/ins
    Burst Credit: 40MiB/s/TiB (max 1TiB credit/TiB)
    Burst Rate: 250MiB/TiB max burst: 500MiB/TiB
    Balanced at 12TiB
    Cold HDD sc1
    500GiB ~ 16TiB, based on 1 MiB/io
    Max Throughput: 250MiB/s/vol, 250iops/vol, 80000iops/ins
    Burst Credit: 12MiB/s/TiB(max 1TiB credit/TiB)  
    Burst Rate: 80MiB/s/TiB max burst: 80MiB/TiB
    Balanced at 20.84TiB(but max 16TiB)
##### Snapshot
    [page](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSSnapshots.html)
    
### Import/Export
    AWS Import/Export does not currently support export from Amazon EBS or Amazon Glacier.
    Note that the maximum supported size for a single object in Amazon S3 is 5TB. Any objects larger than 5TB will not be transferred.
    The amount of data you can transfer is limited by the capacity of the device, up to 16TB

### Cloud Front
    Cloud Front is a global service, and metrics are available only when you choose the US East (N. Virginia) region in the AWS console.

    Cloud Front delivers your content using a global network of edge locations and works seamlessly with Amazon S3 which stores the original and definitive versions of your files. 
    Amazon Cloud Front billing is mainly affected by Data Transfer Out Edge Location Traffic Distribution Requests Dedicated IP SSL Certificates

### Identity and Access Management, IAM
##### Password Policy
    
    [page](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_passwords_account-policy.html)
    IAM account password policy
        You can use a password policy to do these things: 
        
        Set a minimum password length. 
        Require specific character types, including uppercase letters, lowercase letters, numbers, and non-alphanumeric characters. 
        Be sure to remind your users that passwords are case sensitive. 
        Allow all IAM users to change their own passwords. 
        Require IAM users to change their password after a specified period of time (enable password expiration). 
        Prevent IAM users from reusing previous passwords. 
        Force IAM users to contact an account administrator when the user has allowed his or her password to expire.
##### API key
    2 API key/account. API key should be created before access api. use same policy of account to user.
##### Account report
    To access your AWS account resources, users must have credentials. You can generate and download a credential report that lists all users in your account and the status of their various credentials, including passwords, access keys, MFA devices, and signing certificates. You can get a credential report using the AWS Management Console, the AWS CLI, or the IAM API. You can use credential reports to assist in your auditing and compliance efforts. You can use the report to audit the effects of credential lifecycle requirements, such as password rotation. You can provide the report to an external auditor, or grant permissions to an auditor so that h she can download the report directly. You can generate a credential report as often as once every four hours. When you request a report. IAM first checks whether a report for the account has been generated within the past four hours. If so, the most recent report is downloaded. If the most recent report for the account is more than four hours old, or if there are no previous reports for the account, AM generates and downloads a new report. Credential reports are downloaded as comma-separated values (CSV) files. You can open CSV files with common spreadsheet software to perform analysis, or you can build an application that consumes the CSV files programmatically y and performs custom analysis.
##### tmp
    
    An organization has multiple AWS accounts to isolate a development environment from a testing or production environment. At times the users from one account need to access resources in the other account, such as promoting an update from the development environment to the production environment. In this case the IAM role with cross account access will provide a solution. Cross account access lets one account share access to their resources with users in the other AWS accounts.


# NetWork

### Virtual Private Cloud, VPC
##### Capacity
    + VPCs per region, 
        Internet gateways per region, 
        Virtual private gateways per region, 
        Elastic IP addresses per region for EC2-Classic,
        Elastic IP addresses per region for EC2-VPC: 5
    + VPN connections per region: 50
    + VPN connections per VPC: 10
    + Routes per route table: 50
    + Route tables per VPC,
        Subnets per VPC,
        Security groups per VPC: 200
    + NAT gateways per Availability Zone: 5
    + The allowed block size is between a /28 netmask and /16 netmask.
##### default VPC
    a VPC with a size /16 IPv4 CIDR block
    a size /20 default subnet in each Availability Zone
    Create an internet gateway and connect it to your default VPC.
    a main route table for your default VPC with a rule that sends all IPv4 traffic destined for the internet to the internet gateway.
    Create a default security group and associate it with your default VPC.
    Create a default network access control list (ACL) and associate it with your default VPC.
    Associate the default DHCP options set for your AWS account with your default VPC.
##### Subnet
    By design, each subnet must be associated with a network ACL. Every subnet that you create is automatically associated with the VPC's default network ACL.

    type:
    - public subnet: traffic is routed to an internet gateway.
    - private subnet: traffic can't route to the internet gateway.
    - VPN-only subnet: traffic can't route to the internet gateway, but virtual private gateway.

    The first four IP addresses and the last IP address in each subnet CIDR block are not available for you to use, and cannot be assigned to an instance.
##### default Subnet
    You cannot specify the CIDR block yourself.
    You cannot restore a previous default subnet that you deleted.
    Only one default subnet per Availability Zone.
    You cannot create a default subnet in a nondefault VPC.
##### Access Control List, ACL
    Act as a firewall for associated subnets, controlling both inbound and outbound traffic at the subnet level and supports allow rules and deny rules.
    use ACLs to grant basic read/write permissions to AWS accounts and not for public access over the Internet. 
    Evaluate rules from lower number to higher.
##### Default ACL
    default ACL, allow all inbound and outbound
    customized ACL, deny all inbound and outbound
##### VPC Security Groups 
    Act as a firewall for associated Amazon EC2 instances, controlling both inbound and outbound traffic at the instance level and supports allow rules only

    Supports allow rules only, cant deny rules. (Default: allow all inbound and outbound)
##### VPC peering
	- You cannot have overlapping CIDR blocks between VPCs. 
    - You cannot peer VPCs between VPCs in different regions. 
    - You have a limit on the number active and pending VPC peering connections that you can have per VPC. 
    - VPC peering does not support transitive peering relationships; 
    - One VPC peering connection between the same two VPCs at the same time. 
    - The Maximum Transmission Unit (MTU) across a VPC peering connection is 1 500 bytes(封包長度). 
    - A placement group can span peered VPCs; however, you will not get fullbisection bandwidth between instances in peered VPCs. 
    - Unique cast reverse path forwarding in VPC peering connections is not supported. 
    - You cannot reference a security group from the peer VPC as a source or destination for ingress or egress rules in your security group. Instead, reference CIDR blocks of the peer VPC as the source or destination of your security groups ingress or egress rules. 
    - Private DNS values cannot be resolved between instances in peered VPCs.

### Flow log
    You can create a flow log for a VPC, a subnet, or a network interface. If you create a flow log for a subnet or VPC, each network interface in the VPC or subnet is monitored.

    https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html 

    If you launch more instances into your subnet after you've created a flow log for your subnet or VPC, then a new log stream (for CloudWatch Logs) or log file object (for Amazon S3) is created for each new network interface as soon as any network traffic is recorded for that network interface.

    You can create flow logs for network interfaces that are created by other AWS services; for example, Elastic Load Balancing, Amazon RDS, Amazon ElastiCache, Amazon Redshift, and Amazon WorkSpaces. However, you cannot use these service consoles or APIs to create the flow logs; you must use the Amazon EC2 console or the Amazon EC2 API. Similarly, you cannot use the CloudWatch Logs or Amazon S3 consoles or APIs to create flow logs for your network interfaces.

### Elastic Load Balancer, ELB
    - Application Load Balancer, support dynamic host port mapping for ECS.
    - Network Load Balancer
    - Classic Load Balancer

    Access log
##### Cross Zoen Load Balancing

	If the load balancer nodes for your Classic Load Balancer can distribute requests regardless of Availability Zone, this is known as cross-zone load balancing. With cross-zone load balancing enabled, your load balancer nodes distribute incoming requests evenly across the Availability Zones enabled for your load balancer. Otherwise, each load balancer node distributes requests only to instances in its Availability Zone.
##### Connection Draining
    1 ~ 3600 (default: 300sec)
    When you enable connection draining, you can specify a maximum time for the load balancer to keep connections alive before reporting the instance as de-registered. The maximum timeout value can be set between 1 and 3,600 seconds (the default is 300 seconds). When the maximum time limit is reached, the load balancer forcibly closes connections to the de-registering instance.
##### Security Group

### Auto Scaling
    - Amazon EC2 Auto Scaling groups
    - Aurora DB clusters
    - DynamoDB global secondary indexes
    - DynamoDB tables
    - ECS services
    - Spot Fleet requests

##### Life Cycle

    https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroupLifecycle.html
##### EC2 Termination
	- If there are instances in multiple Availability Zones, select the Availability Zone with the most instances and at least one instance that is not protected from scale in. 
    - If there is more than one Availability Zone with this number of instances, select the Availability Zone with the instances that use the oldest launch configuration.
    - Determine which unprotected instances in the selected Availability Zone use the oldest launch configuration.
    - If there are multiple instances that use the oldest launch configuration, determine which unprotected instances are closest to the next billing hour. (This helps you maximize the use of your EC2 instances and manage your Amazon EC2 usage costs.) 
    - If there is more than one unprotected instance closest to the next billing hour, select one of these instances at random.

    Customized termination policies:
    + OldestInstance.
    + NewestInstance.
    + OldestLaunchConfiguration.
    + ClosestToNextInstanceHour.
    + Default.

    Even though the user has configured the termination policy, before Auto Scaling selects an Instance to terminate, It first Identifies the Availability Zone that has more instances than the other Availability Zones used by the group. Within the selected Availability Zone, it identifies the instance that matches the specified termination policy.
##### Limit
	- max autoscalling group: 20
	- Default cooldown period: 300sec
##### Launch Configuration
    configuration for ec2 instance
    launch config is part of auto scalling group


### Direct Connect
	VPN connect between on-premise and aws

	1Gbps and 10Gbps ports are available
    Speeds of 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, and 500Mbps can be ordered from any APN partners supporting AWS Direct Connect.
##### AWS VPN cloud hub
    
    AWS VPN CloudHub leverages an Amazon VPC virtual private gateway with multiple gateways, each using unique BGP autonomous system numbers (ASNs). Your gateways advertise the appropriate routes (BGP prefixes) over their VPN connections. These routing advertisements are received and readvertised to each BGP peer so that each site can send data to and receive data from the other sites. The remote network prefixes for each spoke must have unique ASNs, and the sites must not have overlapping IP ranges. Each site can also send and receive data from the VPC as if they were using a standard VPN connection.

### Route 53

    TLD, top-level domain
##### Route Policy
    - Latency-based routing
    - Weight-based routing
##### Hosted Zone
    A Hosted Zone refers to a selection of resource record sets hosted by Route 53.
	
    In Amazon Route 53, you cannot create a hosted zone for a top-level domain (TLD).
#####Health Check

    To enable DNS Failover for an ELB endpoint, create an Alias record pointing to the ELB and set the “Evaluate Target Health‟ parameter to true. Route 53 creates and manages the health checks for your ELB automatically. You do not need to create your own Route 53 health check of the ELB. You also do not need to associate your resource record set for the ELB with your own health check, because Route 53 automatically associates it with the health checks that Route 53 manages on your behalf. The ELB health check will also inherit the health of your backend instances behind that ELB.
### VPC endpoint
##### Interface Endpoints
	- Amazon API Gateway
	- AWS CloudFormation
	- Amazon CloudWatch
	- Amazon CloudWatch Events
	- Amazon CloudWatch Logs
	- AWS CodeBuild
	- AWS Config
	- Amazon EC2 API
	- Elastic Load Balancing API
	- AWS Key Management Service
	- Amazon Kinesis Data Streams
	- Amazon SageMaker Runtime
	- AWS Secrets Manager
	- AWS Security Token Service
	- AWS Service Catalog
	- Amazon SNS
	- AWS Systems Manager
	- Endpoint services hosted by other AWS accounts
	- Supported AWS Marketplace partner services
##### Gateway Endpoints
    + Amazon S3
    + DynamoDB

### API Gateway
##### Capacity
    - by default, 10000 requests/s. 1000 requests/s at spike.
    - burst limit ??? 5000 requests




# Database

### Relational Database Service, RDS
    - Sychronize automatically in multiple AZ 
    - No Administrative Intervention
        DB Instance failover is fully automatic and requires no administrative intervention. Amazon RDS monitors the health of your primary and standbys, and initiates a failover automatically in response to a variety of failure conditions.
    - Failover conditions
        * Loss of availability in primary Availability Zone
        * Loss of network connectivity to primary
        * Compute unit failure on primary
        * Storage failure on primary
    - Asynchronous for read replic
    - Cross Region Replic should be trigger manually. 
##### Launch an RDS DB in a VPC
	- The VPC must have at least one subnet in at least two of the Availability Zones in the region where you want to deploy your DB instance. For detailed information about creating a VPC with one or more subnets, see the section To create a VPC for use with an Aurora DB cluster in Create a VPC and Subnets.
	- You must create a DB subnet group that contains the VPC and subnet information for the VPC you created. You can create a DB subnet group via the AWS RDS console by selecting Subnet Groups. For more information, see Create an RDS Subnet Group and Working with DB Subnet Groups.
	- If you want your RDS DB instance to be publicly accessible, you must enable the DNS Hostnames and DNS Resolution attributes of your VPC, as described in Updating DNS Support for Your VPC.
##### Capacity
    - Nuber of Cluster: 40
    - DB instance: 40
    - Read replicas per master: 5
    - Total storage: 100TiB  
##### Migration

### Aurora
    ?
    SSD-based

    Amazon Aurora automatically replicates your storage six ways, across three Availability Zones.

    Amazon Aurora storage is fault-tolerant, transparently handling the loss of up to two copies of data without affecting database write availability and up to three copies without affecting read availability. 
    Amazon Aurora storage is also self-healing. 
    Data blocks and disks are continuously scanned for errors and replaced automatically. 

### DynamoDB
	- Dynamo DB Fine-Grained Access Control (FGAC): 
        You can now use AWS Identity and Access Management (IAM) policies to regulate access to items and attributes stored in DynamoDB tables, without the need for a middle-tier proxy as illustrated above.

    - Support atomic counter:

    - Amazon DynamoDB supports fast in-place updates. You can increment or decrement a numeric attribute in a row using a single API call. Similarly, you can atomically add or remove to sets, lists, or maps. View our documentation for more information on atomic updates.

    - All items with the same partition key are stored together, and for composite partition keys, are ordered by the sort key value. DynamoDB will split partitions by sort key if the collection size grows bigger than 10 GB.

    - Read Capacity Unit: 4KB? percentage of throughput for cost
    - Write Cpacity Unit: 1KB? percentage of throughput for cost

    - DynamoDB synchronously replicates data across three facilities in an AWS Region

    [Partition Key](https://aws.amazon.com/blogs/database/choosing-the-right-dynamodb-partition-key/)
##### type
    - string
    - number
    - boolean
##### Primary Key
    - Partition key
    - Combination key:Partition key and sort key
##### Secondary Index
    - Global Secondary Index: partition
        an index with a hash and range key that can be different from those on the table. A global secondary index is considered “global” because queries on the index can span all of the data in a table, across all partitions.
    - Local Secondary Index: sort key
        an index that has the same hash key as the table, but a different range key. A local secondary index is “local” in the sense that every partition of a local secondary index is scoped to a table partition that has the same hash key
##### Consistency Model
    - Eventually consistent reads (the default) – The eventual consistency option maximizes your read throughput. However, an eventually consistent read might not reflect the results of a recently completed write. All copies of data usually reach consistency within a second. Repeating a read after a short time should return the updated data.
    
    - Strongly consistent reads — In addition to eventual consistency, DynamoDB also gives you the flexibility and control to request a strongly consistent read if your application, or an element of your application, requires it. A strongly consistent read returns a result that reflects all writes that received a successful response before the read.
##### Auto Scaling
    1. You create an Application Auto Scaling policy for your DynamoDB table.
    2. DynamoDB publishes consumed capacity metrics to Amazon CloudWatch.
    3. If the table's consumed capacity exceeds your target utilization (or falls below the target) for a specific length of time, Amazon CloudWatch triggers an alarm. You can view the alarm on the AWS Management Console and receive notifications using Amazon Simple Notification Service (Amazon SNS).
    4. The CloudWatch alarm invokes Application Auto Scaling to evaluate your scaling policy.
    5. Application Auto Scaling issues an UpdateTable request to adjust your table's provisioned throughput.
    6. DynamoDB processes the UpdateTable request, dynamically increasing (or decreasing) the table's provisioned throughput capacity so that it approaches your target utilization.

    Decreasing Provisioned Throughput: A decrease is allowed up to four times any time per day. up to 27 times when no descrease in past hour.?

### Simple DB
	
    - AWS simple db similar with AWS DynamoDB without scalibility
##### Capacity
	- 10GB per domain
    - 1b attributes per domain
    - Attribute and item length: 1024 bytes


# Auto Managed Services

### Simple Queue Service, SQS
	Resource-based permissions are supported.

	- A single client can send or receive Amazon SOS messages at a rate of about 5 to 50 messages per second. 

    batch processing,
        Higher receive performance can be achieved by requesting multiple messages (up to 10) in a single call, It may take several seconds before a message that has been to a queue is available to be received.
##### type
    - FIFO queue (ordered). Never duplicate.
    - standard queue (unordered). Delivered at least once.
##### Capacity
    - By default, a message is retained for 4 days. The minimum is 60 seconds (1 minute). The maximum is 1,209,600 seconds (14 days).
    - Standard queues support a nearly unlimited number of transactions per second (TPS) per action.
    - FIFO queues support up to 300 messages per second 
    - The default (minimum) invisibility period for a message is 0 seconds. The maximum is 15 minutes.
    - The minimum message size is 1 byte (1 character). The maximum is 262,144 bytes (256 KB).
    - The default (minimum) visibility timeout for a message is 30 seconds. The maximum is 12 hours.
    - maximum of 120,000 (20,000 for FIFO) inflight messages (received from a queue by a consumer, but not yet deleted from the queue).
##### Dead-Letter Queue
    
    https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html


### Cloud Watch
    - You can‟t use IAM to control access to Cloud Watch data for specific resource. (No CloudWatch Amazon Resource Names (ARNs))
    - cloud watch  metric every 5 mins for 10 intervals?
    - Mornitoring period: basic 5mins, Detailed 1mins
    - metrics only available for 2 weeks.

##### Metric
    Auto Scaling: default 5 mins, extra fee for detailed monitoring
        detailed monitoring is enabled when which create from cli.
    SNS, EMR: default 5 mins only.
    Route 53: default 1 mins without extra fee.
##### Cloud Watch Event
    allow subscription to AWS API calls

    Service list:
	- Amazon EC2 instances
	- AWS Lambda functions
	- Streams in Amazon Kinesis Data Streams
	- Delivery streams in Amazon Kinesis Data Firehose
	- Amazon ECS tasks
	- Systems Manager Run Command
	- Systems Manager Automation
	- AWS Batch jobs
	- Step Functions state machines
	- Pipelines in AWS CodePipeline
	- AWS CodeBuild projects
	- Amazon Inspector assessment templates
	- Amazon SNS topics
	- Amazon SQS queues
	- Built-in targets—EC2 CreateSnapshot API call, EC2 RebootInstances API call, EC2 StopInstances API call, and EC2 TerminateInstances API call.
	- The default event bus of another AWS account
##### Cloud Watch Alarm
    Period, period of time for each datapoint
    Evaluation Period, number of datapoint is required to evaluate
    Datapoints to Alarm, number of datapoint is required to breach
    For example:
        Period: 5, Evaluation Period: 3, Datapoints to Alarm:3 => 15mins with 3 breach
        Period: 5, Evaluation Period: 5, Datapoints to Alarm:4 => 25mins with 4 breach

    Alarm State:
    - OK—The metric is within the defined threshold.
    - ALARM—The metric is outside of the defined threshold.
    - INSUFFICIENT_DATA—The alarm has just started, the metric is not available, or not enough data is available for the metric to determine the alarm state.

    Once CloudWatch retrieves these data points, the following happens:
        If no data points in the evaluation range are missing, CloudWatch evaluates the alarm based on the most recent data points collected.

        If some data points in the evaluation range are missing, but the number of existing data points retrieved is equal to or more than the alarm's Evaluation Periods, CloudWatch evaluates the alarm state based on the most recent existing data points that were successfully retrieved. In this case, the value you set for how to treat missing data is not needed and is ignored.

        If some data points in the evaluation range are missing, and the number of existing data points that were retrieved is lower than the alarm's number of evaluation periods, CloudWatch fills in the missing data points with the result you specified for how to treat missing data, and then evaluates the alarm. However, any real data points in the evaluation range, no matter when they were reported, are included in the evaluation. CloudWatch uses missing data points only as few times as possible.


### Lambda
    Each function associates an execution role (IAM role).
	- For accessing other services from lambda, grant permission of other services to this execution role.
	- For accessing lambda function from other services, grant event source to access lambda function.

    For event sources, except for the stream-based services (Amazon Kinesis Data Streams and DynamoDB) or Amazon SQS queues, you must grant the event source permissions to invoke your AWS Lambda function.

    For poll-based event sources (Amazon Kinesis Data Streams and DynamoDB streams and Amazon SQS queues), AWS Lambda polls the resource on your behalf and reads new records. To enable this, you need to grant AWS Lambda permissions to access the new records. In turn, AWS Lambda will invoke any Lambda function subscribed to this event source to process the event.

    For any other event source that will invoke your Lambda function directly, you must grant that event source permissions to invoke your AWS Lambda function.
##### Capacity
    - max 5mins execution time
    - 100 concurrent requests/s, ask for increment if required.

### Simple Notification Service, SNS

	Resource-based permissions are supported.	
##### Subscriber
    - Lambda
    - SQS
    - SMS
    - Email
    - HTTP/S
##### Publisher
    - Auto Scaling group
    - EC2
    - VPC
    - S3
    - Cloud Watch
##### Capacity
    
    default, SNS offers 10 million subscriptions per topic, and 100,000 topics per account
##### Event Source
    https://docs.aws.amazon.com/lambda/latest/dg/invoking-lambda-function.html
	- Amazon S3
	- Amazon DynamoDB
	- Amazon Kinesis Data Streams
	- Amazon Simple Notification Service
	- Amazon Simple Email Service
	- Amazon Simple Queue Service
	- Amazon Cognito
	- AWS CloudFormation
	- Amazon CloudWatch Logs
	- Amazon CloudWatch Events
	- AWS CodeCommit
	- Scheduled Events (powered by Amazon CloudWatch Events)
	- AWS Config
	- Amazon Alexa
	- Amazon Lex
	- Amazon API Gateway
	- AWS IoT Button
	- Amazon CloudFront
	- Amazon Kinesis Data Firehose
	- Other Event Sources: Invoking a Lambda Function On Demand
	- Sample Events Published by Event Sources


### Simple Email Service, SES

    Amazon SES console--This method is the quickest way to set up your email server.
##### Setup Email Server with SES
    1. Amazon SES console, 
    2. the Simple Mail Transfer Protocol (SMTP) interface, or 
    3. you can call the Amazon SES API. 
##### Send an email with DKIM
    
    Domain Keys Identified Mail (DKIM) is a standard that allows senders to sign their email messages and IPS, and use those signatures to verify that those messages are legitimate and have not been modified by a third party in transit.    
##### Capacity
    - 10,000 emails per 24-hour period
    - maximum rate of 5 emails per second.
    - Every Amazon SES sender has a unique set of sending limits
    - Amazon SES automatically adjusts these limits upward, as long as you send highquality email. 

     If your existing quota is not adequate for your needs and the system has not automatically increased your quota, you can submit an SES Sending Quota Increase case at any time. Sending limits are based on recipients rather than on messages. 

     You can check your sending limits at any time by using the Amazon SES console. 

     Note that if your email is detected to be of poor or questionable quality (e.g., high complaint rates, high bounce rates, spam, or abusive content), Amazon SES might temporarily or permanently reduce your permitted send volume, or take other action as AWS deems appropriate.

### Cloud Trail
    - Cloud trail is not turned on automatically

### Amazon Kinesis
	fully managed serviced

##### type
    - Kinesis Video Stream, Capture, process, and store video streams
    - Kinesis Data Stream, Capture, process, and store data streams
    - Kinesis Data Firehose, Load data streams into AWS data stores
    - Kinesis Data Analytics, Analyze data streams with standard SQL

### Amazon Web Application Firewall, WAF
### Amazon Simple Work Flow, SWF

# Others
### Testing
##### AWS Penetration test
	- User should get approvel from AWS.
	- https://aws.amazon.com/tw/security/penetration-testing/

### Billing
	
    Consolidated Billing enables you to consolidate payment for multiple AWS accounts within your company by designating a single paying account.

### Cloud Hardware Security Module, Cloud HSM
	
    In AWS Cloud HSM, you can perform a remote backup/restore of a Luna SA partition if you have purchased a Luna Backup HSM.

### Amazon Elastic Transcoder: Media Converter






### non relevant
	- OLTP : online transaction processing
	- OLAP : online analytical processing
	- ip sec 
	- internet protocol security
	- Domain name system security extension

	- You can use a mnemonic, such as CPFSS
	- Amazon Resource Names (ARN5) arn: https://docs.aws.amazon.com/zh_tw/general/latest/gr/aws-arns-and-namespaces.html

check list

VPC ACL for subnet
VPC security gruop for ec2

Could subnet config ACL or SG


route 53 hosted zone 
route 53 health check relationship with elb?






?The maximum size of an Amazon EBS snapshot is 1 TB, so if the device image is larger than 1 TB, the image is chunked and stored on Amazon S3. The target location is determined based on the total capacity of the device, not the amount of data on the device
    Snapshot is stored at S3
    Snapshots are incremental backups
    Amazon EBS volumes that serve as root devices, you should stop the instance before taking the snapshot.



    api call
    While using the EC2 GET requests as URLs, the is the URL that serves as the entry point for the web service.?
    The endpoint is the URL that serves as the entry point for the web service.
    When the user account has reached the maximum number of EC2 instances, It will not be allowed to launch an instance. AWS will throw an "Instance Limit Exceeded" error. 
    For all other reasons, such as AMI is missing part, Corrupt Snapshot” or Volume limit has reached” it will launch an EC2 instance and then terminate it.



    - AWS IAM SAML (Security Assertion Markup Language)?
    - AWS IAM Security Token Service (STS)?


    DMS
    ?
    You want to use Oracle Data Dump to import complex databases or databases that are several hundred megabytes or several terabytes in size



auto scalling for network io throughput? how many targets can be used for 
elb health check period? cloud watch 


如果edge location掛了 cloud front會從其他region的拿嗎