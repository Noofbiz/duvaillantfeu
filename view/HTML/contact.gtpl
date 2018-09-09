{{define "contact"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    {{template "head"}}
  </head>
  <body>
    {{template "navbar"}}
    <!-- Page Content -->
    <div class="container">
      <!-- General Contact -->
      <div class="row">
        <div class="panel">
            <div class="panel-body">
              <div class="row">
                <div class="col-lg-12">
                    <h2 class="text-center">Contact Us</h2>
                </div>
              </div>
              <div class="row">
                <div class="col-md-offset-1 col-md-3 col-xs-8 col-xs-offset-2">
                  <script type="text/javascript" id="clstr_globe" src="//cdn.clustrmaps.com/globe.js?d=xnzTakoIYq9rTPat_aZZIkYdpMeJPvhNy9hA9OYt8OM"></script>
                </div>
                <div class="col-md-6 col-md-offset-2">
                    <h5 class="text-center vert_center">
                      Lauren Trathen <br>
                      Du Vaillant Feu Beaucerons <br>
                      <a href="mailto:Vaillantfeubeaucerons@hotmail.com"> Vaillantfeubeaucerons@hotmail.com </a> <br>
                    </h5>
                </div>
              </div>
            </div>
        </div>
      </div> <!-- /General Contact Row -->

      <hr class="watermelon">

      <!-- Social Media -->
      <div class="row">
        <div class="panel">
            <div class="panel-body">
              <div class="col-lg-12">
                  <h2 class="text-center">Find Us on Social Media!</h2>
              </div>
              <div class="col-md-12">
                  <h1 class="text-center">
                    <a href="https://www.facebook.com/vaillantfeubeaucerons"><i class="fa fa-facebook-square" style="color:#3b5998;" aria-hidden="true"></i></a>
                    <a href="https://www.instagram.com/jamaisthebeauceron/"><i class="fa fa-instagram" stlye="color:#125688;" aria-hidden="true"></i></a>
                  </h1>
              </div>
            </div>
        </div>
      </div><!-- /Social Media row -->
    </div>
  <!-- /.container -->

    {{template "jsfoot"}}
  </body>
</html>
{{end}}
