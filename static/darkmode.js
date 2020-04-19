var checkbox = document.getElementById("darkToggle");
window.addEventListener("load", function () {
  if (checkbox) {
    darkmode();
    checkbox.addEventListener("change", function () {
      darkmodeRegister();
    });
  }
});

function darkmode() {
  var darkmodeStat = localStorage.getItem("darkmode") !== null && localStorage.getItem("darkmode") === "dark";

  checkbox.checked = darkmodeStat;
  darkmodeStat ? document.body.setAttribute("data-theme", "dark") : document.body.removeAttribute("data-theme");
}

function darkmodeRegister() {
  if (checkbox.checked) {
    document.body.setAttribute("data-theme", "dark");
    localStorage.setItem("darkmode", "dark");
  } else {
    document.body.removeAttribute("data-theme");
    localStorage.removeItem("darkmode");
  }
}