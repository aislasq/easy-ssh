# essh
Syntactic Sugar for SSH/SFTP connections

### Install
clone this repository locally
<code>go install </code>

### Configuration
<code> nano $HOME/.essh/config.yaml</code>

<pre>
hosts:
  - name: "host1"
    connection:
      hostname: "10.1.1.1"
    credentials:
      user: "username"
      key: "KEY_NAME_1"
  - name: "host2"
    connection:
      hostname: "192.1.1.1"
    credentials:
      user: "new_username"
      key: "KeyName2"
  - name: "host3"
    connection:
      hostname: "50.1.1.1"
      port: 2022
    credentials:
      user: "user_with_password"
keys:
  - name: "KEY_NAME_1"
    path: "/home/my_user/keys/host1_key.pem"
  - name: "KeyName2"
    path: "/home/my_user/keys/my_key_2.pem"
</pre>

### Usage
##### List all configured hosts with their corresponding connections

<code>essh view</code>
<pre>
************************HOSTS***********************
Name   User@Host                          Key       
------+----------------------------------+----------
host1  username@10.1.1.1                 KEY_NAME_1 
host2  new_username@192.1.1.1            KeyName2   
host3  user_with_password@50.1.1.1:2022             

*********************KEYS********************
Name        Path                             
-----------+---------------------------------
KEY_NAME_1  /home/my_user/keys/host1_key.pem 
KeyName2    /home/my_user/keys/my_key_2.pem  
</pre>

##### Connect to a host

<code>essh connect host1</code>
<pre>
Connecting to host1 via username@10.1.1.1

Last login: Sat Aug 31 18:00:00 2019 from 50.100.150.250

       __|  __|_  )
       _|  (     /   Amazon Linux AMI
      ___|\___|___|

https://aws.amazon.com/amazon-linux-ami/2018.03-release-notes/
[username@ip-10-1-1-1 ~]$ 
</pre>

<code>essh sftp host2</code>
<pre>
Connecting SFTP to host2 via new_username@192.1.1.1

Connected to new_username@192.1.1.1.
sftp> 
</pre>

