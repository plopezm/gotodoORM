app.directive('todoCreate', [function(){
	return {
		restrict: "AE",
		controller: "todoController",
		controllerAs: "todoController",
		templateUrl:"js/modules/todo/todo-create/todo-create.html"
	}
}]);
