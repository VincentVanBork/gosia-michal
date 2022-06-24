async function fetch_invitation(token) {
	let response = await fetch("/guest/invitations/get/" + token)
	return await response.json()
}

const showPlacesFn = async function() {
	window.scrollTo(0,0)
	document.body.style.overflow = "hidden"

    let invitation = await fetch_invitation(invitation_token)
    let tableID = invitation.TableId

	const wrapper = document.createElement('div')
	wrapper.setAttribute('id', 'wrapper')
	document.body.appendChild(wrapper)

	const closePlaces = document.createElement('button')
	closePlaces.setAttribute('id', 'closePlaces')
	closePlaces.classList.add('closePlace')
	closePlaces.style.left = '10px'
	closePlaces.innerHTML += 'Wróć do zaproszenia'
	closePlaces.addEventListener('click', closePlacesFn)
	wrapper.appendChild(closePlaces)
	const showPlaces = document.getElementById('showPlaces')
	showPlaces.remove()
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
			if(i==tableID){
			setTimeout(() => {
				const deTable = document.getElementById('table' + tableID)
				chooseTable()
				fade(deTable)
			}, 3000)}
		}
	}, 1000)
}

const closePlacesFn = () => {
	document.body.style.overflow = "auto"
	wrapper.animate([{ transform: 'translateY(100vh)' }, { transform: 'translateY(0vh)' }], {
		duration: 1000,
		iterations: 1,
		fill: 'forwards',
	})
	const closePlaces = document.getElementById('closePlaces')
	closePlaces.remove()

	setTimeout(() => {
		const wrapper = document.getElementById('wrapper')
		wrapper.remove()
		const showPlaces = document.createElement('button')
		showPlaces.setAttribute('id', 'showPlaces')
		showPlaces.classList.add('showPlace')
		showPlaces.innerHTML += 'Pokaż miejsce'
		showPlaces.addEventListener('click', showPlacesFn)
		const imgBox = document.getElementById('img-div-id')
		imgBox.appendChild(showPlaces)
	}, 1000)
}

const showPlaces = document.createElement('button')
showPlaces.setAttribute('id', 'showPlaces')
showPlaces.classList.add('showPlace')
showPlaces.innerHTML += 'Pokaż miejsce'
showPlaces.addEventListener('click', showPlacesFn)
const imgBox = document.getElementById('img-div-id')
imgBox.appendChild(showPlaces)

const drawTable = i => {
	const view_port = window.innerWidth
	const table = document.createElement('div')
	table.setAttribute('class', 'table')
	table.setAttribute('id', 'table' + i)
	if(view_port>1000){
		table.style.transform = 'scale(1)'
	}
	table.innerHTML += i
	wrapper.appendChild(table)
}

const chooseTable = async function() {
    let invitation = await fetch_invitation(invitation_token)
    let tableID = invitation.TableId
	console.log(tableID);
	const tableCh = document.getElementById('table' + tableID)
	tableCh.style.color = 'rgb(9, 9, 169)'
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
	table.style.color = 'rgba(100, 100, 100, 0.5)'
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
