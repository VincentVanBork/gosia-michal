async function fetch_invitation(token) {
    let response = await fetch("/guest/invitations/get/" + token)
    return await response.json()
}

const queryString = window.location.search;
console.log(queryString);
const urlParams = new URLSearchParams(queryString);
const sl = urlParams.get('sl');
const invitation_token = urlParams.get("token")
const audio_mix = [
    "1.mp3",
    "2.mp3",
    "3.mp3"
];

function get_random(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

var txt = [
    'Zaproszenie dla',
    (sl == 1 ? "Ciebie" : "Was"),
    ["Z radością zapraszamy " + (sl == 1 ? "Cię" : "Was")],
    "na uroczyste zawarcie sakramentu małżeństwa,",
    "które odbędzie się 25 czerwca 2022 roku o godzinie 17:00.",
    ["Liczymy na " + (sl == 1 ? "Twoją" : "Waszą") + " obecność podczas wspólnego świętowania!"]
];

var speed = 97;
let inx = 0;

const typer = (txt, box, inx) => {
    if (inx < txt.length) {
        document.getElementById(box).innerHTML += txt.charAt(inx);
        inx++;
        setTimeout(() => typer(txt, box, inx), speed);
        //console.log(inx, txt)
    }
}

function unfade(element, disp) {
    var op = 0.1;
    element.style.display = disp;
    var timer = setInterval(function () {
        if (op >= 1) {
            clearInterval(timer);
        }
        element.style.opacity = op;
        element.style.filter = 'alpha(opacity=' + op * 100 + ")";
        op += op * 0.1;
    }, 10);
}

document.addEventListener('DOMContentLoaded', async function (event) {
    let src_audio = audio_mix[get_random(0, 2)];
    let audio = new Audio("mp3/" + src_audio);
    let img_box = document.getElementById("img-div-id")
    img_box.addEventListener('click', function (event) {
        audio.play();
    })
    document.addEventListener('swiped', function (e) {
        audio.play();
    });
    let invitation = await fetch_invitation(invitation_token)
    console.log(invitation.Guests)
    if (invitation.Guests.length > 0) {
        txt[1] = ""
    }
    for (let i = 0; i < invitation.Guests.length; i++) {
        txt[1] += " "
        txt[1] += invitation.Guests[i].FirstName
        txt[1] += " "
        txt[1] += invitation.Guests[i].LastName
        if (i === 0 && invitation.Guests.length > 1) {
            txt[1] += " "
            txt[1] += "i"
        }
        if (!invitation.IsSingle) {
            if (i === 0 && invitation.Guests.length === 1) {
                txt[1] += " "
                txt[1] += "wraz z osobą towarzysząca"
            }
        }
    }
    setTimeout(() => typer(txt[0], "f1t", inx), 100);

    setTimeout(() => typer(txt[1], "f2t", inx), speed * txt[0].length);

    setTimeout(() => typer(txt[2][0], "f1tb", inx), speed * (txt[1].length + txt[0].length));
    setTimeout(() => typer(txt[3], "f2tb", inx), speed * (txt[2][0].length + txt[1].length + txt[0].length));
    setTimeout(() => typer(txt[4], "f3tb", inx), speed * (txt[3].length + txt[2][0].length + txt[1].length + txt[0].length));
    setTimeout(() => typer(txt[5][0], "f4tb", inx), speed * (txt[4].length + txt[3].length + txt[2][0].length + txt[1].length + txt[0].length));

    let element1 = document.querySelector(".bottom-box");
    setTimeout(() => unfade(element1, "flex"), speed * (txt[5][0].length + txt[4].length + txt[3].length + txt[2][0].length + txt[1].length + txt[0].length));

    let element2 = document.querySelector(".reg-box");
    setTimeout(() => unfade(element2, "block"), speed * (txt[5][0].length + txt[4].length + txt[3].length + txt[2][0].length + txt[1].length + txt[0].length));
});

console.log(window.location.pathname.split("/").at(-1))