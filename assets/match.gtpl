<html>
    <head>
    <title>Match</title>
      <style>
      table, th, td {
    	   border: 1px solid black;
    	   border-collapse: collapse;
	   }
      </style>
    </head>
    <body>
      <form action="/match" method="post">
        <input type="submit" name="Choice" value="Rock">
        <input type="submit" name="Choice" value="Paper">
        <input type="submit" name="Choice" value="Scissors">
      </form>
      <p>{{.Result}}<p>
      <br>
      <table style="width:100%">
      <tr>
      <th>Firstname</th>
      <th>Lastname</th>
      <th>Age</th>
      </tr>
      </table>
    </body>
</html>
