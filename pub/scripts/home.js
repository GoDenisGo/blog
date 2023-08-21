// el holds all our handlers
const el = {};

function prepareHandles() {
  el.cv = document.querySelector('.cv');
  el.blog = document.querySelector('.blog');
}

function pageLoaded() {
  prepareHandles();
}

pageLoaded();
