<!DOCTYPE html>
<html>
<head>
    <style>
        body { margin: 0; }
        #application-content { display: none; }
    </style>
</head>
<body>
    <div id="application-layout"></div>
    <div id="application-content">
        <h1>Legacy codebase</h1>
    </div>
    <script>
        window.__APP_LAYOUT = {
            getLayoutTarget: function() {
                return document.getElementById('application-layout');
            },
            getContentTarget: function() {
                return document.getElementById('application-content');
            },
        };
    </script>
    <script src="/layout/shell.js" type="module"></script>
</body>
</html>
