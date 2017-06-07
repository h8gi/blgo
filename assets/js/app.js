var MyVue = Vue.extend({
  delimiters: ['((', '))'],
  methods: {	
  }
})

var fragmentsList = new MyVue({
  el: '#fragments-list',
  data: {
	fragments: []
  },
  methods: {
	update: function(e) {
	  axios.get('/api/fragments')
		.then(function(response){
		  fragmentsList.fragments = response.data
		})
		.catch(function(error){
		  console.log(error)
		})
	},
	deleteFragment: function(name) {
	  axios.delete('/api/fragments/' + name)
		.then(function(response){
		  fragmentsList.update()
		})
		.catch(function(error){
		  console.log(error)
		})
	}
  }
})

fragmentsList.update()

var editor = new MyVue({
  el: '#editor',
  data: {
	name: '',
    contents: ''
  },
  methods: {
    post: function() {
      payload = new FormData()
      payload.append('contents', this.contents)
	  payload.append('name', this.name)
      axios.post('/api/fragments', payload)
        .then(function(response){
          fragmentsList.update()
        })
        .catch(function(error){
          console.log(error)
        })
    }
  }
})
