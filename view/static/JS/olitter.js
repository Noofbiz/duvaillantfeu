var images = {
  'nova': [
    "/static/images/olitter2018/nova/0.jpg",
    "/static/images/olitter2018/nova/1.jpg",
    "/static/images/olitter2018/nova/2.jpg",
  ],
  'ocean':
  [
    "/static/images/olitter2018/ocean/0.jpg",
    "/static/images/olitter2018/ocean/1.jpg",
    "/static/images/olitter2018/ocean/2.jpg",
    "/static/images/olitter2018/ocean/3.jpg",
  ],
  'oiseau': [
    "/static/images/olitter2018/oiseau/0.jpg",
    "/static/images/olitter2018/oiseau/1.jpg",
    "/static/images/olitter2018/oiseau/2.jpg",
    "/static/images/olitter2018/oiseau/3.jpg",
  ],
  'olivier': [
    "/static/images/olitter2018/olivier/0.jpg",
    "/static/images/olitter2018/olivier/1.jpg",
    "/static/images/olitter2018/olivier/2.jpg",
  ],
  'omen': [
    "/static/images/olitter2018/omen/0.jpg",
    "/static/images/olitter2018/omen/1.jpg",
    "/static/images/olitter2018/omen/2.jpg",
    "/static/images/olitter2018/omen/3.jpg",
  ],
  'ozark': [
    "/static/images/olitter2018/ozark/0.jpg",
    "/static/images/olitter2018/ozark/1.jpg",
    "/static/images/olitter2018/ozark/2.jpg",
    "/static/images/olitter2018/ozark/3.jpg",
    "/static/images/olitter2018/ozark/4.jpg",
    "/static/images/olitter2018/ozark/5.jpg",
  ],
  'roux': [
    "/static/images/olitter2018/roux/0.jpg",
    "/static/images/olitter2018/roux/1.jpg",
    "/static/images/olitter2018/roux/2.jpg",
    "/static/images/olitter2018/roux/3.jpg",
  ],
  'sinatra': [
    "/static/images/olitter2018/sinatra/0.jpg",
    "/static/images/olitter2018/sinatra/1.jpg",
  ],
  'zinc': [
    "/static/images/olitter2018/zinc/0.jpg",
    "/static/images/olitter2018/zinc/1.jpg",
    "/static/images/olitter2018/zinc/2.jpg",
    "/static/images/olitter2018/zinc/3.jpg",
    "/static/images/olitter2018/zinc/4.jpg",
    "/static/images/olitter2018/zinc/5.jpg",
    "/static/images/olitter2018/zinc/6.jpg",
    "/static/images/olitter2018/zinc/7.jpg",
    "/static/images/olitter2018/zinc/8.jpg",
    "/static/images/olitter2018/zinc/9.jpg",
    "/static/images/olitter2018/zinc/10.jpg",
    "/static/images/olitter2018/zinc/11.jpg",
    "/static/images/olitter2018/zinc/12.jpg",
    "/static/images/olitter2018/zinc/13.jpg",
    "/static/images/olitter2018/zinc/14.jpg",
    "/static/images/olitter2018/zinc/15.jpg",
  ],
};

var addedAlready = {
  'nova': [],
  'ocean': [],
  'oiseau': [],
  'olivier': [],
  'omen': [],
  'ozark': [],
  'roux': [],
  'sinatra': [],
  'zinc': [],
};

$( document ).ready(function() {
  $(".slideshow").each(function(){
    var numToAdd = getImageToAdd(this);
    $(this).attr("src", images[getExtraClass(this)][numToAdd]);
    window.setTimeout(imageMagic, rand(2000, 5000), this);
  });
});

function getExtraClass(ele) {
  var extraClass = '';
  if ($(ele).hasClass('nova')) {
    extraClass = 'nova';
  } else if ($(ele).hasClass('ocean')) {
    extraClass = 'ocean';
  } else if ($(ele).hasClass('oiseau')) {
    extraClass = 'oiseau';
  } else if ($(ele).hasClass('olivier')) {
    extraClass = 'olivier';
  } else if ($(ele).hasClass('omen')) {
    extraClass = 'omen';
  } else if ($(ele).hasClass('ozark')) {
    extraClass = 'ozark';
  } else if ($(ele).hasClass('roux')) {
    extraClass = 'roux';
  } else if ($(ele).hasClass('sinatra')) {
    extraClass = 'sinatra';
  } else if ($(ele).hasClass('zinc')) {
    extraClass = 'zinc';
  }
  return extraClass
}

function placeNewBehind(ele, numToAdd){
  var extraClass = getExtraClass(ele)
  var imgStr = '<img class="img-responsive img-portfolio slideshow ';
  imgStr += extraClass + '" src="'
  imgStr += images[extraClass][numToAdd] + '" style="display:none;"';
  imgStr += '>';
  var parent = $(ele).parent();
  parent.prepend(imgStr);
  return $(ele).parent().children('img:nth-child(1)');
}

function imageMagic(ele){
  // begin switching out old image
  $(ele).removeClass('slideshow_active');
  $(ele).addClass('slideshow_switching');
  var newImage = placeNewBehind(ele, getImageToAdd(ele));

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

function getImageToAdd(ele){
  var extraClass = getExtraClass(ele)
  if(addedAlready[extraClass].length == images[extraClass].length){
    addedAlready[extraClass] = [];
  }
  while(true){
    var num = rand(0, images[extraClass].length - 1);
    if(!addedAlready[extraClass].includes(num)){
      addedAlready[extraClass].push(num);
      return num;
    }
  }
}
