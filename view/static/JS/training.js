var images = ["/static/images/training/slideshow/1.jpeg",
"/static/images/training/slideshow/2.jpg",
"/static/images/training/slideshow/3.jpeg",
"/static/images/training/slideshow/4.jpg",
"/static/images/training/slideshow/5.jpg",
"/static/images/training/slideshow/6.jpg",
"/static/images/training/slideshow/8.jpeg",
"/static/images/training/slideshow/9.jpeg",
"/static/images/training/slideshow/10.jpeg",
"/static/images/training/slideshow/11.JPG",
"/static/images/training/slideshow/12.JPG",
"/static/images/training/slideshow/13.jpeg",
"/static/images/training/slideshow/14.jpeg",
"/static/images/training/slideshow/15.JPG",
"/static/images/training/slideshow/16.JPG",
"/static/images/training/slideshow/17.JPG",
"/static/images/training/slideshow/18.JPG",
"/static/images/training/slideshow/19.jpeg",
"/static/images/training/slideshow/20.jpeg",
"/static/images/training/slideshow/22.jpeg",
"/static/images/training/slideshow/24.jpg",
"/static/images/training/slideshow/25.jpg",
"/static/images/training/slideshow/26.jpg",
"/static/images/training/slideshow/27.jpeg",
"/static/images/training/slideshow/about.jpg"];

var addedAlready = [];

$( document ).ready(function() {
  $(".slideshow").each(function(){
    var numToAdd = getImageToAdd();
    $(this).attr("src", images[numToAdd]);
    window.setTimeout(imageMagic, rand(2000, 5000), this);
  });
});

function placeNewBehind(ele, numToAdd){
  var imgStr = '<img class="img-responsive img-portfolio slideshow" src="';
  imgStr += images[numToAdd] + '" style="display:none;"';
  imgStr += '">'
  $(ele).parent().append(imgStr);
  return $(ele).parent().children('img:nth-child(2)')
}

function imageMagic(ele){
  // begin switching out old image
  $(ele).removeClass('slideshow_active');
  $(ele).addClass('slideshow_switching');
  var newImage = placeNewBehind(ele, getImageToAdd());

  //fadeout old image
  $(ele).fadeOut(1000, function(){
    $(ele).remove();
    $(newImage).fadeIn(1000, function(){
      $(newImage).addClass('slideshow_active');
      window.setTimeout(imageMagic, rand(2000, 5000), newImage);
    });
  });
}

function rand(min, max){
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getImageToAdd(){
  if(addedAlready.length == images.length){
    addedAlready = [];
  }
  while(true){
    var num = rand(0, images.length - 1);
    if(!addedAlready.includes(num)){
      addedAlready.push(num);
      return num;
    }
  }
}
