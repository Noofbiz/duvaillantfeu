{{define "navbar"}}
<!-- Navigation -->
 <nav class="navbar navbar-inverse navbar-fixed-top" role="navigation" id="navi">
     <div class="container">
         <!-- Brand and toggle get grouped for better mobile display -->
         <div class="navbar-header">
             <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                 <span class="sr-only">Toggle navigation</span>
                 <span class="icon-bar"></span>
                 <span class="icon-bar"></span>
                 <span class="icon-bar"></span>
             </button>
             <a class="navbar-brand" href="/">Beaucerons du Vaillant Feu</a>
         </div>
         <!-- Collect the nav links, forms, and other content for toggling -->
         <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
             <ul class="nav navbar-nav navbar-right">
                 <li>
                     <a href="/BreedStandard">Breed Standard</a>
                 </li>
                 <li class="dropdown">
                     <a href="#" class="dropdown-toggle" data-toggle="dropdown">Litters<b class="caret"></b></a>
                     <ul class="dropdown-menu">
                         <li>
                           <a href="/Litters/Upcoming">Upcoming Litters</a>
                         </li>
                         <hr>
                         <li>
                          <a href="/Litters/OLitter2018">O-Litter Spring 2018</a>
                         </li>
                         <!--
                         <hr>
                         <li>
                             <a href="/Litters/Care">Caring for a Beauceron</a>
                         </li>
                       -->
                     </ul>
                 </li>
                 <li class="dropdown">
                     <a href="#" class="dropdown-toggle" data-toggle="dropdown">Meet Our Dogs<b class="caret"></b></a>
                     <ul class="dropdown-menu">
                         <li>
                             <a href="/Parents/Jamais">Jamais</a>
                         </li>
                         <li>
                             <a href="/Parents/Mowgli">Mowgli</a>
                         </li>
                         <li>
                             <a href="/Parents/Oshi">Océane</a>
                         </li>
                         <li>
                             <a href="/Parents/Gateau">Gâteau</a>
                         </li>
                     </ul>
                 </li>
                 <li>
                     <a href="/Apply">Puppy Application</a>
                 </li>
                 <li>
                     <a href="/FAQ">FAQs</a>
                 </li>
                 <li>
                     <a href="/Contact">Contact Us</a>
                 </li>
             </ul>
         </div>
         <!-- /.navbar-collapse -->
     </div>
     <!-- /.container -->
 </nav>
{{end}}
