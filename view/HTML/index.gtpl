{{define "index"}}
<!DOCTYPE html>
<html lang="en">
  <head profile="http://www.w3.org/2005/10/profile">
    {{template "head"}}
  </head>
  <body>

    <div id="fb-root"></div>
    <script>(function(d, s, id) {
      var js, fjs = d.getElementsByTagName(s)[0];
      if (d.getElementById(id)) return;
      js = d.createElement(s); js.id = id;
      js.src = "//connect.facebook.net/en_US/sdk.js#xfbml=1&version=v2.10&appId=142519669677356";
      fjs.parentNode.insertBefore(js, fjs);
      }(document, 'script', 'facebook-jssdk'));
    </script>

    {{template "navbar"}}

    <!-- Header Carousel -->
    <header id="myCarousel" class="carousel slide">
      <!-- Wrapper for slides -->
      <div class="carousel-inner">
        <div class="item active">
            <div class="fill" style="background-image:url('/static/images/index/Slide1.jpeg');"></div>
        </div>
      </div>
    </header>

    <!-- Page Content -->
    <div class="container">

      <!-- Welcome to Site -->
      <div class="row">
        <div class="panel">
            <div class="panel-body">
              <div class="col-md-6 col-md-offset-3">
                  <img class="img-responsive img-rounded" src="/static/images/index/VFBSlogan2.png" alt="Beaucerons du Vaillant Feu 'Difficult to describe, unmistakable when present'">
              </div>
              <div class="col-md-4">
                  <p> Welcome to the Vaillant Feu kennel, a small volume preservation breeder in Southern California focused on breeding the best in versatility, stability, health and temperament. As a preservation breeder, it is our job to preserve and protect the legacy of the Beauceron and work to move the breed forward. All dogs being bred in this kennel have received Excellent ratings at our Journée de Beauceron from French judges and have proven themselves in different working capacities. Here at Vaillant Feu, we pride ourselves on our beaucerons being as true to working form as they can be and being the "country gentleman" that our standard calls for; never timid or harsh, always unwaveringly frank and courageous. </p>
                  <p> We are a small up and coming breeder located in San Diego County, California. Our passion for the breed is represented in the dogs we choose to breed and the work we dedicate them to. At Vaillant Feu, conformation to the standard with an aesthetically pleasing functional form, stability in all environments and the necessary drives to compete are what we put first. The Beauceron is an ancient French symbol of partnership and dedication to the job. We hope to produce for our fellow breed enthusiasts a dog you can do anything with, a versatile beauceron with exceptional temperament. </p>
              </div>
              <div class="col-md-4">
                  <img class="img-responsive img-rounded" src="/static/images/index/jamaiswork.jpg" alt="Lauren (the breeder) standing next to Jamais.">
              </div>
              <div class="col-md-4">
                <!-- Facebook Social Plugin! -->
                <div class="fb-page" data-href="https://www.facebook.com/vaillantfeubeaucerons" data-tabs="timeline" data-small-header="true" data-adapt-container-width="true" data-hide-cover="true" data-show-facepile="false" data-width=500><blockquote cite="https://www.facebook.com/MixieMelts/" class="fb-xfbml-parse-ignore"><a href="https://www.facebook.com/MixieMelts/">MixieMelts</a></blockquote></div>
                <!-- /Facebook -->
              </div>
            </div>
        </div>
      </div>
      <!-- /.row -->

      <hr class="watermelon">

      <!-- About Beauceron -->
      <div class="row">
        <div class="panel">
          <div class="panel-body">
            <div class="row">
              <div class="col-lg-12">
                  <h2 class="page-header text-right">The Beauceron</h2>
              </div>
            </div>
            <div class="row">
              <div class="col-md-6">
                <div class="match-box">
                  <iframe src="https://www.youtube.com/embed/PLsiBOfcIRU?ecver=1" allowfullscreen class="framezor"></iframe>
                </div>
              </div>
              <div class="col-md-6">
                <div class="match-box">
                  <p> A protector, solemnly guarding his charges; be it family or flock. The Beauceron is a confident, highly intelligent, independent and strong-willed dog. An active and intelligent breed, they require plenty of mental and physical exercise.</p>
                  <p>The Beauceron is a versatile breed that has been selected and bred for a very long time, first mentioned in a manuscript written in 1587! The Beauceron have been used for many types of work due to their intelligence and work ethic. Today's Beauceron is being used much in the same fashion as the German Shepherd dog in this country. While still tending sheep and cattle, the Beauceron is utilized for military and police work, Search & Rescue, handicapped assistance work, canine sports such as Agility, French Ring, Schutzhund, Obedience, Tracking, Skijoring, and of course as a family companion.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- /.row -->
  <!-- Page hit counter! -->
  <div id="phc">
    <a href="https://clustrmaps.com/site/19ihr" title="Visit tracker">
      <img src="https://www.clustrmaps.com/map_v2.png?d=xnzTakoIYq9rTPat_aZZIkYdpMeJPvhNy9hA9OYt8OM&cl=ffffff">
    </a>
  </div>
  <!-- /.counter -->

    {{template "jsfoot"}}
    <script src="//cdnjs.cloudflare.com/ajax/libs/jquery.matchHeight/0.7.0/jquery.matchHeight-min.js"></script>
    <script>
      $(document).ready(function(){
        $('.match-box').matchHeight();
      });
    </script>
  </body>
</html>
{{end}}
