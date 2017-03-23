<html>
	<head>
  		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	</head>
	<body onload="onload()">
		<form action="/login" method="post">
			Name:<input type="text" name="name"/><br/>
			PassWord:<input type="password" name="password"/><br/>
			<input type="submit" value="submit"/>
		</form>
		<script type="text/javascript">
			function onload(){
				var msg = {{.msg}};
				if(msg)
					alert(msg);
			}
		</script>
	</body>
</html>