#get argument from command line
import getdbcon
import argparse
import pyperclip


parser = argparse.ArgumentParser(description='qlip: quick clipboard')
parser.add_argument('-a', '--attached', help='attached your clipboard to the cloud', action='store_true')
parser.add_argument('-p', '--paste', help='print the lastes clipboard paste', action='store_true')
parser.add_argument('-t', '--tabel', help='show tabel', action='store_true')
parser.add_argument('-k', '--key', help='add a key to your clip', action='append')


args = parser.parse_args()



#check if the user wants to attach the clipboard to the cloud
if args.attached:
    if args.key:
        print(f"Keys: {args.key}")
        #get the clipboard content
        content = pyperclip.paste()
        #get the database connection
        cnx = getdbcon.getdbcon()
        #insert the content into the database
        getdbcon.insert_key_and_paste(cnx, content, args.key[0])
    else:
        #get the clipboard content
        content = pyperclip.paste()
        #get the database connection
        cnx = getdbcon.getdbcon()
        #insert the content into the database
        getdbcon.insert(cnx, content)


if args.paste:
    print(args.attached)
    #get the database connection
    cnx = getdbcon.getdbcon()
    #get the latest clipboard content
    pyperclip.copy(getdbcon.getlatest(cnx))
    #close the connection
    getdbcon.closecnx(cnx)
    #paste the content

if args.tabel:
    #get the database connection
    cnx = getdbcon.getdbcon()
    #get the latest clipboard content
    getdbcon.showtabel(cnx)

if args.key and not args.attached:
    #get the database connection
    cnx = getdbcon.getdbcon()
    #get the latest clipboard content
    getdbcon.getpastebykey(cnx, args.key[0])
    #close the connection
    getdbcon.closecnx(cnx)


