$( document ).ready(function() {
  $('#toggle-doggle').change(function() {
    var unhide;
    var hide1;
    var hide2;
    var hideall = false;

    switch($('#toggle-doggle option:selected').text()){
      case "Household Pet":
        unhide = ".household-pet";
        hide1 = ".sport-quality";
        hide2 = ".show-quality";
        break;
      case "Sport/Working Quality":
        unhide = ".sport-quality";
        hide1 = ".household-pet";
        hide2 = ".show-quality";
        break;
      case "Show Quality":
        unhide = ".show-quality";
        hide1 = ".household-pet";
        hide2 = ".sport-quality";
        break;
      case "":
        hideall = true;
        break;
    }
    if (hideall) {
      $(".show-quality").addClass("hidden");
      $(".sport-quality").addClass("hidden");
      $(".household-pet").addClass("hidden");
    } else {
      $(unhide).removeClass("hidden");
      $(hide1).addClass("hidden");
      $(hide2).addClass("hidden");
    }
    hideall = false;
  });
});
