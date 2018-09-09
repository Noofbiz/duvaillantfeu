{{define "apply_success"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    {{template "head"}}
  </head>
  <body>
    {{template "navbar"}}
    <!-- Page Content -->
    <div class="container">
      <div class="panel panel-success">
        <div class="panel-heading">
          <div class="row">
            <div class="col-xs-8 col-xs-offset-2">
              <h3 class="text-center">
                Application successfully sent! <br>
                We will get back to you shortly after reviewing your application. <br>
              </h3>
            </div>
          </div>
        </div>
        <div class="panel-body">
          <p class="text-center">Click <a href="/">here</a> to return to the homepage.</p>
        </div>
      </div>
    </div>
  <!-- /.container -->

    {{template "jsfoot"}}
  </body>
</html>
{{end}}
