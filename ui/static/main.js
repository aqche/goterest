// Navbar Burger Toggling
const navbarBurger = document.querySelector(".navbar-burger");
const navbarMenu = document.querySelector(".navbar-menu");

navbarBurger.addEventListener("click", () => {
  navbarBurger.classList.toggle("is-active");
  navbarMenu.classList.toggle("is-active");
});

// Notification Deletion
const deleteButtons = document.querySelectorAll(".delete");
deleteButtons.forEach(deleteButton => {
  deleteButton.addEventListener("click", () => {
    deleteButton.parentNode.remove();
  });
});

// Masonry Setup
let masonry;

window.addEventListener("load", () => {
  let grid = document.querySelector(".grid");
  if (grid) {
    masonry = new Masonry(grid, {
      itemSelector: ".grid-item",
      gutter: 10,
      fitWidth: true,
      columnWidth: 238
    });
  }
});

window.addEventListener("resize", () => {
  if (masonry) {
    masonry.layout();
  }
});

// Image Error Handling
function handleImgErr(img) {
  img.onerror = undefined;
  img.src = "https://via.placeholder.com/238x238?text=Image+Not+Found";
}
