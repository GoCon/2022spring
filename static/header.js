window.addEventListener('scroll', function(){
  if (this.window.scrollY >= 100) {
    document.getElementById('navigation-menu').className = 'menu menu--white';
    document.getElementById('header').className = 'header--white';
    document.getElementById('logo-text').className = 'logo-text logo--disable';
    document.getElementById('logo-text--white').className = 'logo-text--white';
  } else {
    document.getElementById('navigation-menu').className = 'menu';
    document.getElementById('header').className = '';
    document.getElementById('logo-text').className = 'logo-text';
    document.getElementById('logo-text--white').className = 'logo-text--white logo--disable';
  }
});
