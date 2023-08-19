// el holds all our handlers
const el = {};

function prepareHandles() {
  el.cv = document.querySelector('.cv');
  el.blog = document.querySelector('.blog');
}

function btnCVListener() {
  const payload = {
    resource: 'cv',
  };
  el.cv.addEventListener('click', async () => {
    console.log('Clicked CV button.');
    await fetchCV(payload);
  });
}

async function fetchCV(payload) {
  const res = await fetch('/' + payload.resource, {
    method: 'POST',
    headers: { 'Content-type': 'application/json' },
    body: JSON.stringify(payload),
  });

  if (res.ok) {
    console.log('Payload success.');
  } else {
    console.log('Payload fail.');
  }
}

function btnBlogListener() {
  const payload = {
    resource: 'blog',
  };
  el.blog.addEventListener('click', async () => {
    console.log('Clicked Blog button.');
    await fetchBlog(payload);
  });
}

async function fetchBlog(payload) {
  const res = await fetch('/' + payload.resource, {
    method: 'POST',
    headers: { 'Content-type': 'application/json' },
    body: JSON.stringify(payload),
  });

  if (res.ok) {
    console.log('Payload success.');
  } else {
    console.log('Payload fail.');
  }
}

// prepareListeners adds an event listener to each element in el
function prepareListeners() {
  btnCVListener();
  btnBlogListener();
}

function pageLoaded() {
  console.log('home page loaded'); // development info, temporary
  prepareHandles();
  prepareListeners();
}

pageLoaded();
