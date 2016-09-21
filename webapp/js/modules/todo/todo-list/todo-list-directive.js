app.directive('todoList', [function(){
	return {
		restrict: "AE",
		controller: "todoController",
		controllerAs: "todoController",
		templateUrl:"js/modules/todo/todo-list/todo-list.html"
	}
}]);
