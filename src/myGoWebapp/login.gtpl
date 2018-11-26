<html>
    <head>
    <title>User Operations</title>
    </head>
    <body>
        <form action="/create" method="post">
            Username:<input type="text" name="username">
            Department:<input type="text" name="department">
            Date Created:<input type="text" name="datecreated">
            <input type="submit" value="Create User">
        </form>

        <form action="/update" method="post">
            Username:<input type="text" name="username">
            Department:<input type="text" name="department">
            <input type="submit" value="Update User">
        </form>

        <form action="/delete" method="post">
            Username:<input type="text" name="username">
            
            <input type="submit" value="Delete User">
        </form>

    </body>
</html>
