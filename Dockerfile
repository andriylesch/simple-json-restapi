# our image will be based from alpine image 

##########################################################
# syntax for image versions 
# FROM alpine         - will download latest version of alpine image from hub.docker.com (https://hub.docker.com/_/alpine/)
# FROM alpine:3.1     - will download alpine version 3.1 - of course if you need it

FROM alpine

##########################################################
# MAINTAINER <YOUR FIRSTNAME AND LASTNAME> <YOUR_EMAIL_ADDRESS> - This is a non-executable instruction used to indicate the author of the Dockerfile.

MAINTAINER Andriy Lesch clasikas@gmail.com

##########################################################
# RUN <Command name>  This instruction let's you execute a command on top of an existing layer and create a new layer with the results of command execution. 

# BAD PRACTICE
# RUN apt-get update
# RUN apt-get install -y golang
# RUN apt-get install -y nginx
# RUN etc ..... 

# GOOD PRACTICE (will reduce number of layers)  
# RUN apt-get update && apt-get install -y golang && apt-get install -y nginx && etc ....

# in current article we don't need it

##########################################################
# EXPOSE - While running your service in the container you may want your container to listen on specified ports. The EXPOSE instruction helps you do this.
# EXPOSE <PORT_NUMBER>

EXPOSE 5000

##########################################################
# ENV - This instruction can be used to set the environment variables in the container.

# FOR EXAMPLE
# ENV var_home="/tmp/etc" 


##########################################################
# COPY - used to copy files and directories from a specified source to a destination (in the file system of the container).

# FOR EXAMPLE
# COPY temp.txt /tmp/etc
# COPY <YOUR FILE/DIRECTORIES> to <DIRECTORY OF CONTAINER>  

##########################################################
# ADD - the the same as COPY instruction with few added features like remote URL support in the source field and local-only tar extraction. 
# But if you don’t need a extra features, it is suggested to use COPY as it is more readable.

# FOR EXAMPLE
# ADD http://your_url.com/some.tar.xz /tmp/usr
# ADD URL to <DIRECTORY OF CONTAINER>


##########################################################
# ENTRYPOINT - it can be used to set the primary command for the image.

