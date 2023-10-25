#get argument from command line
import sys
import pyautogui
import os
import getdbcon
import argparse
import pyperclip


#parse argument
parser = argparse.ArgumentParser(description='qlip: quick clipboard')
parser.add_argument('-a', '--attached', help='attached your clipboard to the cloud', action='store_true')
parser.add_argument('-p', '--paste', help='print the lastes clipboard paste', action='store_true')
parser.add_argument('-t', '--tabel', help='show tabel', action='store_true')

args = parser.parse_args()


#check if the user wants to attach the clipboard to the cloud
if args.attached:
    #get the clipboard content
    content = pyperclip.paste()
    #get the database connection
    cnx = getdbcon.getdbcon()
    #insert the content into the database
    getdbcon.insert(cnx, content)



if args.paste:
    #get the database connection
    cnx = getdbcon.getdbcon()
    #get the latest clipboard content
    pyperclip.copy(getdbcon.getlatest(cnx))
    #close the connection
    getdbcon.closecnx(cnx)
    #paste the content
    #smiolate a left mouse click
    pyautogui.click()
    pyautogui.hotkey('ctrl', 'v')

if args.tabel:
    #get the database connection
    cnx = getdbcon.getdbcon()
    #get the latest clipboard content
    getdbcon.showtabel(cnx)





