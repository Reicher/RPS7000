<html>
    <head>
    <title>Practice</title>
    </head>
    <body>
      <form action="/practice" method="post">
        <input type="submit" name="Choice" value="Rock">
        <input type="submit" name="Choice" value="Paper">
        <input type="submit" name="Choice" value="Scissors">
      </form>
      <p>{{.Result}}<p>
    </body>
</html>
