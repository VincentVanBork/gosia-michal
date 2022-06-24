async function fetch_invitation() {
	let username = "michaljuras"
	let password = "P@ndaExpress"
	// let headers = new Headers();
	// headers.set('Authorization', 'Basic ' + username + ":" + password)
	// headers: headers,
	let response = await fetch("/lista/dane", {method: 'GET', credentials: 'include'})
	return await response.json()
}
var goscie;
window.onload = async ()=> {
	goscie = await fetch_invitation()

	for (let i = 0; i < goscie.length; i++) {
		let person = document.createElement('div')
		person.setAttribute('id', i)
		person.classList.add('person')
		person.innerHTML += goscie[i].Name + ' ' + goscie[i].Surname
		person.addEventListener('click', function () {
			showPlacesFn(goscie[i].TableId)
		})
		document.body.appendChild(person)
	}
}

// const goscie = [
// 	{ id: 0, name: 'Michał', lastname: 'Juraś', tableID: '0' },
// 	{ id: 1, name: 'Michał', lastname: 'Darowny', tableID: '3' },
// 	{ id: 2, name: 'Marta', lastname: 'Juraś', tableID: '2' },
// 	{ id: 3, name: 'Małgorzata', lastname: 'Juraś', tableID: '6' },
// ]

const closePlacesFn = () => {
	wrapper.animate([{ transform: 'translateY(100vh)' }, { transform: 'translateY(0vh)' }], {
		duration: 1000,
		iterations: 1,
		fill: 'forwards',
	})
	const closePlaces = document.getElementById('closePlaces')
	closePlaces.remove()
	// const showPlaces = document.createElement('button')
	// showPlaces.setAttribute('id', 'showPlaces')
	// showPlaces.innerHTML += 'Pokaż miejsce'
	// showPlaces.addEventListener('click', showPlacesFn)
	// document.body.appendChild(showPlaces)
	setTimeout(() => {
		//const wrapper = document.getElementById('wrapper')
		//wrapper.remove()
		location.reload()
	}, 1000)
}

const wrapper = document.createElement('div')
wrapper.setAttribute('id', 'wrapper')
document.body.appendChild(wrapper)

const wrapper2 = document.createElement('div')
wrapper2.setAttribute('id', 'wrapper2')
wrapper.appendChild(wrapper2)

const showPlacesFn = tableID => {
	const closePlaces = document.createElement('button')
	closePlaces.setAttribute('id', 'closePlaces')
	closePlaces.innerHTML += 'Wróć do listy'
	closePlaces.addEventListener('click', closePlacesFn)
	wrapper.appendChild(closePlaces)
	// const showPlaces = document.getElementById('showPlaces')
	// showPlaces.remove()
	wrapper.animate([{ transform: 'translateY(0vh)' }, { transform: 'translateY(100vh)' }], {
		duration: 1000,
		iterations: 1,
		fill: 'forwards',
	})
	setTimeout(() => {
		for (let i = 0; i < 9; i++) {
			setTimeout(() => {
				drawTable(i)
				fade(document.getElementById('table' + i))
			}, 250 * (i + 1))

			setTimeout(() => {
				chooseTable(tableID)
				fade(document.getElementById('table' + tableID))
			}, 3000)
		}
	}, 1000)
}



// const showPlaces = document.createElement('button')
// showPlaces.setAttribute('id', 'showPlaces')
// showPlaces.innerHTML += 'Pokaż miejsce'
// showPlaces.addEventListener('click', showPlacesFn)
// document.body.appendChild(showPlaces)

const drawTable = i => {
	const viewport_width = window.innerWidth
	console.log(viewport_width)
	const table = document.createElement('div')
	table.setAttribute('class', 'table')
	table.setAttribute('id', 'table' + i)
	table.innerHTML += i
	wrapper2.appendChild(table)
}

const chooseTable = tableID => {
	const tableCh = document.getElementById('table' + tableID)
	tableCh.style.color = 'gold'
	tableCh.style.fontSize = '100px'
	tableCh.style.opacity = '0'
	tableCh.classList.add('circle__box')

	const cwr = document.createElement('div')
	cwr.classList.add('circle__wrapper', 'circle__wrapper--right')
	tableCh.appendChild(cwr)

	const cwwr = document.createElement('div')
	cwwr.classList.add('circle__whole', 'circle__right')
	cwr.appendChild(cwwr)

	const cwl = document.createElement('div')
	cwl.classList.add('circle__wrapper', 'circle__wrapper--left')
	tableCh.appendChild(cwl)

	const cwwl = document.createElement('div')
	cwwl.classList.add('circle__whole', 'circle__left')
	cwl.appendChild(cwwl)

	const table = document.createElement('div')
	table.setAttribute('class', 'tableABS')
	table.innerHTML += tableID
	table.style.color = 'gold'
	tableCh.appendChild(table)
}

const fade = element => {
	var op = 0.1
	element.style.display = 'block'
	var timer = setInterval(function () {
		if (op >= 1) {
			clearInterval(timer)
		}
		element.style.opacity = op
		element.style.filter = 'alpha(opacity=' + op * 100 + ')'
		op += op * 0.1
	}, 30)
}
