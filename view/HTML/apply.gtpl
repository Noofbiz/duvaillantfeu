{{define "apply"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    {{template "head"}}
    <script src='https://www.google.com/recaptcha/api.js'></script>
  </head>
  <body>
    {{template "navbar"}}
    <!-- Page Content -->
    <div class="container">
      <!-- Application Form  -->
      <div class="panel">
        <div class="panel-body">
          <form action="/Apply/Send">
            <div class="form-group">
              <label>Name</label>
              <input type="text" class="form-control" name="Name">
            </div>
            <div class="form-group">
              <label>Age</label>
              <input type="text" class="form-control" name="Age">
            </div>
            <div class="form-group">
              <label>Phone</label>
              <input type="text" class="form-control" name="Phone">
            </div>
            <div class="form-group">
              <label>Email</label>
              <input type="email" class="form-control" name="Email">
            </div>
            <div class="form-group">
              <label>Address</label>
              <textarea class="form-control" name="Address"></textarea>
            </div>
            <div class="form-group">
              <label>Do you rent or own?</label>
              <select class="form-control" name="RentOrOwn">
                <option value=""></option>
                <option value="Rent">Rent</option>
                <option value="Own">Own</option>
              </select>
            </div>
            <div class="form-group">
              <label>How many members of your household are there, and how old are they?</label>
              <textarea class="form-control" name="Household"></textarea>
            </div>
            <div class="form-group">
              <label>Why do you want a beauceron and what is your experience with the breed?</label>
              <textarea class="form-control" name="WantExp"></textarea>
            </div>
            <div class="form-group">
              <label>What dogs/animals have you owned in the past? How long did you have them and what happened to them if they are no longer in your care?</label>
              <textarea class="form-control" name="PastAnimals"></textarea>
            </div>
            <div class="form-group">
              <label>Vet References: Name and Contact Information.</label>
              <textarea class="form-control" name="VetReference"></textarea>
            </div>
            <div class="form-group">
              <label>Trainer References: Name and Contact Information.</label>
              <textarea class="form-control" name="TrainerReference"></textarea>
            </div>
            <div class="form-group">
              <label>What is your dog training experience?</label>
              <textarea class="form-control" name="DogExp"></textarea>
            </div>
            <div class="form-group">
              <label>What titles and training have you done with your past/current dogs?</label>
              <textarea class="form-control" name="TitlesTraining"></textarea>
            </div>
            <div class="form-group">
              <label>What are you looking for in a puppy?</label>
              <select class="form-control" name="WhatLookingFor" id="toggle-doggle">
                <option value=""></option>
                <option value="Household Pet">Household Pet</option>
                <option value="Sport/Working Quality">Sport/Working Quality</option>
                <option value="Show Quality">Show Quality</option>
              </select>
            </div>
            <div class="form-group household-pet hidden">
              <label>What is your activity level during an average week?</label>
              <input type="text" class="form-control" name="ActivityLevel">
            </div>
            <div class="form-group household-pet hidden">
              <label>How much time can you realistically devote to training your new beauceron puppy indoors and socially out in the world?</label>
              <input type="text" class="form-control" name="RealisticallyTraining">
            </div>
            <div class="form-group household-pet hidden">
              <label>What activities do you plan to use to stimulate your beauceron puppy mentally?</label>
              <input type="text" class="form-control" name="MentalStimulation">
            </div>
            <div class="form-group household-pet hidden">
              <label>What activities do you plan to use to stimulate your beauceron puppy physically?</label>
              <input type="text" class="form-control" name="PhysicalStimulation">
            </div>
            <div class="form-group sport-quality hidden">
              <label>What sport or work do you want a beauceron for?</label>
              <input type="text" class="form-control" name="WhatSport">
            </div>
            <div class="form-group sport-quality hidden">
              <label>Do you have a dedicated club to assist you?</label>
              <input type="text" class="form-control" name="Club">
            </div>
            <div class="form-group sport-quality hidden">
              <label>What level are you hoping to achieve in your sport (if applicable)?</label>
              <input type="text" class="form-control" name="SportLevel">
            </div>
            <div class="form-group sport-quality hidden">
              <label>What are you looking for in terms of drives and energy level?</label>
              <input type="text" class="form-control" name="DriveEnergy">
            </div>
            <div class="form-group show-quality hidden">
              <label>What venues do you plan to show in (AKC/UKC/IABCA)?</label>
              <input type="text" class="form-control" name="Venues">
            </div>
            <div class="form-group show-quality hidden">
              <label>Have you ever titled a dog to Championship or GCH before?</label>
              <input type="text" class="form-control" name="Titled">
            </div>
            <div class="g-recaptcha" data-sitekey="6Lcx6h4UAAAAAF1N5UnvJaFBd_CUdkl3ZEow7QA4"></div>
            <button type="submit" class="btn btn-default">Send Application</button>
          </form>
        </div>
      </div> <!-- /Application Form -->

    </div>
  <!-- /.container -->

    {{template "jsfoot"}}
    <script src="/static/JS/apply.js"></script>
    <!--script src="/static/JS/jqBootstrapValidation.js"></script-->
  </body>
</html>
{{end}}
