this is not a repo to build 1500 account bot upvote network, because for every account you need a CLIENT_ID and CLIENT_SECRET of that account from https://www.reddit.com/prefs/apps

setup:

Download https://github.com/canercetin-randomguy/reddit-upvote/releases/download/doot/reddit-upvote.exe, create a file called proxy.env, credentials.env and post.env

Fill proxy.env like this:
PROXY_1=http://117.160.132.37:9091
PROXY_2=http://121.22.53.166:9091
PROXY_3=http://144.91.66.128:3128
PROXY_4=http://183.239.61.167:9091

just follow the pattern of PROXY_X=????

Fill credentials.env like this
-CLIENT_ID=****
-CLIENT_SECRET=****
-USERNAMEREDDIT=****
-PASSWORD=****
-CLIENT_ID_1=****
-CLIENT_SECRET_1=****
-USERNAMEREDDIT_1=****
-PASSWORD_1=****

You need to grab credentials for every account that in network. https://www.reddit.com/prefs/apps grab it here.

Fill post.env like this:
POST_ID=*****

Simply id of the upvote post.

put everything in the same folder with exe, fire up the exe.

This is why it isnt for 1000+ account bot networks, it is just hassle, with guarantee to work. 
