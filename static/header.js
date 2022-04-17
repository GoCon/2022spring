window.addEventListener("scroll", function () {
  if (this.window.scrollY >= 100) {
    document.getElementById("navigation-menu").className = "menu menu--white";
    document.getElementById("header").className = "header--white";
    document.getElementById("logo-text").className =
      "logo-text logo--black logo--disable";
    document.getElementById("logo-text--white").className =
      "logo-text logo--white";
  } else {
    document.getElementById("navigation-menu").className = "menu";
    document.getElementById("header").className = "";
    document.getElementById("logo-text").className = "logo-text logo--black";
    document.getElementById("logo-text--white").className =
      "logo-text logo--white logo--disable";
  }
});
