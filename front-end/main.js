const ul = document.querySelector("#list-ul");
ul.className = "list";

window.onload = init();
document.getElementById("addBtn").addEventListener("click", toogleModal);
document.getElementById("clr").addEventListener("click", clearFields);
document.getElementById("sv").addEventListener("click", saveBookData)

function createListItem(id, name, author, publisher, isbn) {
    const li = document.createElement("li"),
        p = document.createElement("p");
    li.className = "item";
    p.className = "title";

    // Delete button
    const deleteButton = document.createElement('button');
    deleteButton.addEventListener("click", () => deleteBook(deleteButton.value, ul, li));
    deleteButton.className = "delete btn";
    deleteButton.value = id;
    const deleteIcon = document.createElement("i");
    deleteIcon.className = "bx bx-trash icon";
    deleteButton.appendChild(deleteIcon);

    // Info button
    const infoButton = document.createElement("button");
    infoButton.className = "info btn";
    infoButton.addEventListener("click", () => checkInfo(Number(infoButton.value)));
    infoButton.value = id;
    const infoIcon = document.createElement("i");
    infoIcon.className = "bx bx-info-circle icon";
    infoButton.appendChild(infoIcon);

    p.innerHTML = name;
    li.appendChild(p);
    li.appendChild(deleteButton);
    li.appendChild(infoButton);
    ul.appendChild(li);
}

function init() {
    GetBooksList();
}

function GetBooksList() {
    let listOfBook = [];
    fetch("http://127.0.0.1:8088/api/v1/list")
        .then(response => {
            return response.json()
        })
        .then((list) => {
            list.forEach(book => {
                createListItem(book.ID, book.name, book.author, book.publisher, book.isbn)
                listOfBook.push(book)
            })
            localStorage.setItem("listOfBook", JSON.stringify(listOfBook))
        })
}

function checkInfo(bookId) {
    const infoModal = document.getElementById("infoModal");
    const span = document.getElementById("close-info");

    infoModal.style.display = "block";

    const rawList = localStorage.getItem("listOfBook"),
        list = JSON.parse(rawList);

    const book = list.find(item => item.ID === bookId);

    document.getElementById("bookNameOut").innerHTML = book.name;
    document.getElementById("authorOut").innerHTML = book.author;
    document.getElementById("publisherOut").innerHTML = book.publisher;
    document.getElementById("ISBNOut").innerHTML = book.ISBN;

    span.addEventListener("click", () => {
        infoModal.style.display = "none";
    });

    window.addEventListener("click", (event) => {
        if (event.target === "modal") {
            infoModal.style.display = "none";
        }
    });
}

function deleteBook(bookId, list, item) {
    console.log(bookId);
    list.removeChild(item);
}

function toogleModal() {
    const addModal = document.getElementById("addModal");
    const span = document.getElementById("close");

    addModal.style.display = "block";

    span.addEventListener("click", () => {
        addModal.style.display = "none";
    });

    window.addEventListener("click", (event) => {
        if (event.target === "modal") {
            addModal.style.display = "none";
        }
    });
}

function clearFields() {
    const fields = document.querySelectorAll(".content-input");
    fields.forEach(field => field.value = "");
}

function saveBookData() {
    let obj = {
        Id: Math.floor(Math.random() * 10),
        Name: document.getElementById("bookName").value,
        // Author: document.getElementById("author").value,
        // Publisher: document.getElementById("publisher").value,
        // ISBN: document.getElementById("ISBN").value
    };
    document.getElementById("clr").click();

    createListItem(obj.Id, obj.Name);
}