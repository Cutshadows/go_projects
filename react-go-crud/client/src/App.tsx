import React from 'react'

const App = () => {
  return (
	<div>
		<button onClick={async ()=> {
		const response = await fetch("/users")
		const data = await response.json()
		console.log(data);
		}}>
			Obtener datos
		</button>

	</div>
  )
}

export default App;
