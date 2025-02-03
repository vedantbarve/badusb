# Disclamer

This project is for educational purpose only. Do not execute this attack on anyone without their consent.

# BadUSB

In this world full of cyberattacks, humans are the weakest links to break into a system. The USB rubber ducky or BadUSB is a malicious tool, which when inserted into a system will inject keyboard strokes as if it is an HID (Human Interface Device like keyboard, mouse, game controller etc.).

In this project, I have used Attiny85 as the microcontroller of choice as it is cheap, compact and versatile. This project builds the BadUSB from ground up i.e we will be making the USB, flashing the bootloader, writing a script that runs when we insert the pendrive into the victim's machine and the payload that runs at the client side which compromises the system. This project initiates a reverse shell connection from the victim to the attacker to gain command line access of the victim's machine.

## Aim of the project

1. Run a powershell script on a windows machine after inserting the USB rubber ducky.
2. Initiate a reverse shell connection from the victim to the attacker to gain command line access of the victim's machine.

## Contents

| Sr. No. | Title           | Description                                                                                             |
| ------- | --------------- | ------------------------------------------------------------------------------------------------------- |
| 1.      | badusb-script   | The script that runs when the USB is inserted into the victim's machine.                                |
| 2.      | badusb-programs | The payloads written in GO that establish the reverse-shell connection from the victim to the attacker. |

## Motivation

In an ordinary cyberattack, an attacker tries to find vulnerabilities in the victim's machine by monitoring them over the network. With newer IDS and IPS techniques, it is quite difficult to penetrate a system when an attacker tries to connect with the victim.  
On the other hand, the attacker can make changes on its machine to listen and accept for incoming connections. So if we make the victim connect to the attacker, there are no restrictions by the firewall.  
Thus, a reverse shell attack seemed perfect for this usecase.

## Methodology

![methodology](https://github.com/vedantbarve/BadUSB/blob/master/assets/methodology-1.JPG)
** Socket server port and HTTP server port are changed to :9999 and :9998 respectively, while migrating from C++ to GO.

After the USB rubber ducky is inserted into the victim's machine, following steps take place:

1. Keyboard stokes open powershell and use the curl command to download the payload (client.exe) from the "/" endpoint on the HTTP server hosted on PORT=9998 on the attacker's IP address.
2. We run the client.exe file.
3. On running the client.exe file, victim's machine initiates a connection with the attacker who is listening for incoming connections on its PORT=9999.
4. Once the attacker accepts the incoming connection, we can start exchanging data.

## Softwares used

1. ArduinoIDE :
   - To flash bootloader onto the Attiny85.
   - To burn the script that runs after inserting the USB.
2. KiCAD :
   - To design the PCB for the USB.
3. VScode :
   - To make the GO payload.

## Future scope

This project can be made dynamic by adding payloads for operating systems other than windows. We can send the OS data from the victim to the attacker first, so that the HTTP server will respond with appropriate payload.

## Images

![image-1](https://github.com/vedantbarve/BadUSB/blob/master/assets/504d2812-3e31-429c-a163-e281aa5cedb9.jpeg)

[video-1](https://github.com/vedantbarve/BadUSB/blob/master/assets/bb3333ab-b4e2-4a25-af28-9a1edc2237c5.mov)
## Contact me

Please feel free to contact me via the email given below for any doubts. I am always open for any sort of constructive criticism.

Email :- barvevedant@gmail.com
