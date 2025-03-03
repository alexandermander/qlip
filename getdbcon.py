#creaete a mysql connection to the database
import mysql.connector
from mysql.connector import errorcode

def getdbcon():
    try:
        cnx = mysql.connector.connect(user='', password='', host='azure.com')
        return cnx
    except mysql.connector.Error as err:
        if err.errno == errorcode.ER_ACCESS_DENIED_ERROR:
            print("Something is wrong with the user name or password")
        elif err.errno == errorcode.ER_BAD_DB_ERROR:
            print("Database does not exist")
        else:
            print(err)

def closecnx(cnx):
    cnx.close()

#create a select statement
def select(cnx):
    cnx.database = 'qclip'
    cursor = cnx.cursor()
    query = ("SELECT * FROM pastetb")
    cursor.execute(query)
    for (col1, col2) in cursor:
        print("{} {}".format(col1, col2))
    cursor.close()


#insert a clipboard content into the database
def insert(cnx, content):
    cnx.database = 'qclip'
    cursor = cnx.cursor()
    query = ("INSERT INTO pastetb (paste) VALUES (%s)")
    cursor.execute(query, (content,))
    cnx.commit()
    cursor.close()
#this get the lates clipboard content from the database "SELECT * FROM pastetb WHERE id = (SELECT MAX(id) FROM pastetb);"
def getlatest(cnx):
    cnx.database = 'qclip'
    cursor = cnx.cursor()
    query = ("SELECT * FROM pastetb WHERE id = (SELECT MAX(id) FROM pastetb);")
    cursor.execute(query)
    for (col1, col2, col3) in cursor:
        return col2

def getpastebykey(cnx, key):
    cnx.database = 'qclip'
    cursor = cnx.cursor()
    query = ("SELECT * FROM pastetb WHERE short_key = %s")
    cursor.execute(query, (key,))
    for (col1, col2, col3) in cursor:
        print("{}".format(col2))
        return col2
    cursor.close()

def showtabel(cnx):
    cnx.database = 'qclip'
    cursor = cnx.cursor()
    query = ("SELECT * FROM pastetb")
    cursor.execute(query)
    for (col1, col2, col3) in cursor:
        print("{}".format(col2))
    cursor.close()

def insert_key_and_paste(cnx, content, key):
    cnx.database = 'qclip'
    cursor = cnx.cursor()
    query = ("INSERT INTO pastetb (paste, short_key) VALUES (%s, %s)")
    values = (content, key)
    cursor.execute(query, values)
    cnx.commit()
    cursor.close()

