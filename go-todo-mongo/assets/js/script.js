const taskInput = document.querySelector(".task-input input"),
filters = document.querySelectorAll(".filters span"),
clearAll = document.querySelector(".clear-btn"),
taskBox = document.querySelector(".task-box");

let editId,isEditTask = false

fetchTodos().then(data => showTodo("all",data));
allTodos = "";
filters.forEach(btn => {
    btn.addEventListener("click", () => {
        document.querySelector("span.active").classList.remove("active");
        btn.classList.add("active");
        showTodo(btn.id,"",false);
    });
});

function showTodo(filter,todos = "",firstTime) {
    let todoArr = todos == "" ? allTodos : todos;
    if(firstTime) {
        allTodos = todoArr;
    }
    let liTag = "";
    if(todoArr) {
        todoArr.forEach((todo) => {
            let completed = todo.status == "completed" ? "checked" : "";
            if(filter == todo.status || filter == "all") {
                liTag += `<li class="task">
                            <label for="${todo["ID"]}">
                                <input onclick="updateStatus(this)" type="checkbox" id="${todo["ID"]}" ${completed}>
                                <p class="${completed}">${todo.name}</p>
                            </label>
                            <div class="settings">
                                <i onclick="showMenu(this)" class="uil uil-ellipsis-h"></i>
                                <ul class="task-menu">
                                    <li onclick='editTask(${todo["ID"]}, "${todo.name}")'><i class="uil uil-pen"></i>Edit</li>
                                    <li onclick='deleteTask(${todo["ID"]}, "${filter}")'><i class="uil uil-trash"></i>Delete</li>
                                </ul>
                            </div>
                        </li>`;
            }
        });
    }
    taskBox.innerHTML = liTag || `<span>You don't have any task here</span>`;
    let checkTask = taskBox.querySelectorAll(".task");
    !checkTask.length ? clearAll.classList.remove("active") : clearAll.classList.add("active");
    taskBox.offsetHeight >= 300 ? taskBox.classList.add("overflow") : taskBox.classList.remove("overflow");
}

function showMenu(selectedTask) {
    let menuDiv = selectedTask.parentElement.lastElementChild;
    menuDiv.classList.add("show");
    document.addEventListener("click", e => {
        if(e.target.tagName != "I" || e.target != selectedTask) {
            menuDiv.classList.remove("show");
        }
    });
}

function updateStatus(selectedTask) {
    let taskName = selectedTask.parentElement.lastElementChild;
    let newStatus = "";
    if(selectedTask.checked) {
        taskName.classList.add("checked");
        newStatus = "completed";
        allTodos[selectedTask.id].status = "completed";

    } else {
        taskName.classList.remove("checked");
        newStatus = "pending";
        allTodos[selectedTask.id].status = "pending";
    }
    updateTodo(selectedTask.id,taskName,newStatus).then(data => console.log(data));

}

function editTask(taskId, textName) {
    editId = taskId;
    isEditTask = true;
    taskInput.value = textName;
    taskInput.focus();
    taskInput.classList.add("active");
    updateTodo(taskId,textName).then(data => showTodo("all",data));
}

function deleteTask(deleteId, filter) {
    isEditTask = false;
    allTodos.splice(deleteId, 1);
    deleteTodos(deleteId).then(data => showTodo(filter,data)); // fix
}

clearAll.addEventListener("click", () => {
    isEditTask = false;
    ClearAllTodos().then(data => console.log(data));
    showTodo() // fix
});

taskInput.addEventListener("keyup", e => {
    let userTask = taskInput.value.trim();
    if(e.key == "Enter" && userTask) {
        if(!isEditTask) {
            let taskInfo = {name: userTask, status: "pending"};
            addTodo(taskInfo).then(data => console.log(data));
        } else {
            isEditTask = false;
            todos[editId].name = userTask; // fix
        }
        taskInput.value = "";
        showTodo(document.querySelector("span.active").id); // fix
    }
});


async function ClearAllTodos() {

    const response = await fetch('http://localhost:8080/todos', {
        method: 'DELETE',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          }
    });
    const todos = await response.json();
    return todos;
}

async function updateTodo(id,name,status) {
    const response = await fetch('http://localhost:8080/todo/' + id , {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body: JSON.stringify({
            'name' : name,
            'status' : status
        }
        )
    });
    const todos = await response.json();
    return todos;
}

async function addTodo(todo) { 
    const response = await fetch('http://localhost:8080/todo', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body: JSON.stringify({
            'name' : todo["name"],
            'status' : todo["status"]
        }
        )
    });
   const todos = await response.json();
   return todos;
}

async function fetchTodos() {
    const response = await fetch('http://localhost:8080/todos');
    const todos = await response.json();
    return todos;
}

async function deleteTodos(id) {
    const response = await fetch('http://localhost:8080/todo/' + id, {
        method: 'DELETE'
    });
    const todos = await response.json();
    return todos;
}