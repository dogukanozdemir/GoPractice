const taskInput = document.querySelector(".task-input input"),
filters = document.querySelectorAll(".filters span"),
clearAll = document.querySelector(".clear-btn"),
taskBox = document.querySelector(".task-box");

let editId,isEditTask,editStatus = false
fetchTodos().then(data => showTodo("all",data,true));
allTodos = "";
filters.forEach(btn => {
    btn.addEventListener("click", () => {
        document.querySelector("span.active").classList.remove("active");
        btn.classList.add("active");
        showTodo(btn.id,"",false);
    });
});

function showTodo(filter,todos = "",changeAllTodos) {
    if(changeAllTodos) {
        allTodos = todos;
    }
    let liTag = "";
    if(allTodos) {
        allTodos.forEach((todo) => {
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
                                    <li onclick='editTask("${todo["ID"]}","${todo["name"]}","${todo["status"]}")'><i class="uil uil-pen"></i>Edit</li>
                                    <li onclick='deleteTask("${todo["ID"]}", "${filter}")'><i class="uil uil-trash"></i>Delete</li>
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
        findAndEditTodo(selectedTask.id,taskName.innerText,newStatus);

    } else {
        taskName.classList.remove("checked");
        newStatus = "pending";
        findAndEditTodo(selectedTask.id,taskName.innerText,newStatus);
    }
    updateTodo(selectedTask.id,taskName.innerText,newStatus).then(data => console.log(data));

}

function editTask(taskId, textName,taskStatus) {
    editId = taskId;
    editStatus = taskStatus;
    isEditTask = true;
    taskInput.value = textName;
    taskInput.focus();
    taskInput.classList.add("active");
}

function deleteTask(deleteId, filter) {
    isEditTask = false;
    findAndDeleteTodo(deleteId);
    deleteTodos(deleteId).then(data => {
        showTodo(filter,"",false)
        console.log(data);
    });
}

clearAll.addEventListener("click", () => {
    isEditTask = false;
    allTodos.splice(0, allTodos.length);
    ClearAllTodos().then(data => console.log(data));
    showTodo("all","",false);
});

taskInput.addEventListener("keyup", e => {
    let userTask = taskInput.value.trim();
    if(e.key == "Enter" && userTask) {
        if(!isEditTask) {
            allTodos = !allTodos ? [] : allTodos;
            let taskInfo = {name: userTask, status: "pending"};
            addTodo(taskInfo).then(data => {
                taskInfo["ID"] = data["insertedId"];
                allTodos.push(taskInfo);
                showTodo(document.querySelector("span.active").id,"",false);
                console.log(data);

            });
        } else {
            isEditTask = false;
            updateTodo(editId,userTask,editStatus).then(data => console.log(data));
            findAndEditTodo(editId,userTask,editStatus);
            showTodo(document.querySelector("span.active").id,"",false);
        }
        taskInput.value = "";
    }
});


function findAndDeleteTodo(id) {
    allTodos.forEach((todo,index) => {
        if(todo.ID == id) {
            allTodos.splice(index,1);
        }
    });
    
}
function findAndEditTodo(id,name,status) {
    allTodos.forEach((todo) => {
        if(todo.ID == id) {
            todo.name = name;
            todo.status = status;
        }
    });
}


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
    const response = await fetch('http://localhost:8080/todo/' , {
        method: 'PUT',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
        body: JSON.stringify({
            'ID': id,
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