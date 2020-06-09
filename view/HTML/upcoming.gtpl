{{define "upcoming"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    {{template "head"}}
  </head>
  <body>
    {{template "navbar"}}
    <!-- Page Content -->
    <div class="container">
      <!-- About Jamais -->
      <div class="row">
        <div class="panel">
          <div class="panel-body">
            <div class="col-md-12">
              <h4 class="page-header text-center">There are no upcoming litters at this time.</h4>
            </div>
            <div class="col-md-12">
              <img class="img-responsive img-rounded center-block" src="/static/images/training/slideshow/9.jpeg" alt="Jamais standing on a shopping cart.">
            </div>
          </div>
        </div>
      </div><!-- /About Jamais row -->
    </div>
  <!-- /.container -->

    {{template "jsfoot"}}
  </body>
</html>
{{end}}
