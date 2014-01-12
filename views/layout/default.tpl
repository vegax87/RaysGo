<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="{{.Description}}">
    <meta name="keywords" content="{{.Keywords}}">
    <meta name="author" content="{{.Author}}">

    <title>{{.Title}}</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/bootstrap/css/bootstrap.min.css" rel="stylesheet"/>
    <link href="/static/css/main.css" rel="stylesheet"/>

    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/bootstrap/js/bootstrap.min.js"></script>

    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
    <script src="/static/js/html5shiv.js"></script>
    <script src="/static/js/respond.min.js"></script>
    <![endif]-->
</head>

<body>

{{template "layout/main_nav.tpl" .}}

<div id="main-wrapper" class="container">
    <div id="message" class="container">
        <!-- TODO -->
    </div>
    <div class="clearfix"></div>
    <div id="content">
        {{.LayoutContent}}
    </div>
</div>

{{template "layout/footer.tpl" .}}

</body>
</html>
