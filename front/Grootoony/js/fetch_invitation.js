function fetch_invitation(token) {
    fetch("/api/guest/invitations/get/" + token)
        .then(r => r.json())
        .then(data => console.log(data))
    // return {"names": names, "guestNum": guests, "hasKids": hasKids, "hasWeddingReception": hasFullInvite}
}

export {fetch_invitation}