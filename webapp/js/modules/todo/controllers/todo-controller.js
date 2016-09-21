app.controller('todoController', ['$scope', 'todoService', function($scope, todoService){

	this.todoList = [];
	this.newTodo = {
		name: "",
		desc: ""
	}

	this.init = function(){
		todoService.getTodoList().then(function(response){
			$scope.todoController.todoList = response.data;
		});
	}	

	this.saveTodo = function(){
		todoService.saveTodo(this.newTodo).then(function(response){
			console.log(response);
			if(response.status == 201){
				$scope.todoController.init();
				closeSaveWindow();
			}
		});
	}

	this.removeTodo = function(todo){
		todoService.removeTodo(todo).then(function(response){
			console.log(response);
			if(response.status == 200){
				$scope.todoController.init();
			}
		});
	}
	
	this.init();


	/**
	* JS func
	**/
	function closeSaveWindow(){
		document.getElementById('modalCreate').style.display='block';
	}
}]);

