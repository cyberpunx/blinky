import { HOUSES } from './constants.js';

export function getUserColorByHouse(house){
    switch(house){
        case HOUSES.GRYFFINDOR.NAME:
            return HOUSES.GRYFFINDOR.COLOR;
        case HOUSES.HUFFLEPUFF.NAME:
            return HOUSES.HUFFLEPUFF.COLOR;
        case HOUSES.RAVENCLAW.NAME:
            return HOUSES.RAVENCLAW.COLOR;
        case HOUSES.SLYTHERIN.NAME:
            return HOUSES.SLYTHERIN.COLOR;
        case HOUSES.EXTRANJEROS.NAME:
            return HOUSES.EXTRANJEROS.COLOR;
        default:
            return HOUSES.EXTRANJEROS.COLOR;
    }
}

export function pickRandomColor(){
    const colors = ["primary", "secondary", "success", "danger", "warning", "info", "light", "blue", "azure", "indigo", "purple", "pink", "red", "orange", "yellow", "lime","green", "teal", "cyan", "facebook", "twitter", "google", "vimeo", "linkedin", "dribble", "instagram" ,"pinterest", "vk", "rss", "flickr", "bitbucket", "tabler"];
    return colors[Math.floor(Math.random() * colors.length)];

}
