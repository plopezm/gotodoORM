app.service('todoService',['$http', function($http){
	
	this.getTodoList = function(){
		return $http.get('/api/v1/todos');
	}	

	this.saveTodo = function(newTodo){
		return $http.put('/api/v1/todos', newTodo);
	}

	this.removeTodo = function(todo){
		return $http["delete"]('/api/v1/todos/'+todo.id);
	}

}]);
