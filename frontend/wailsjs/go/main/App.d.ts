// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {config} from '../models';
import {tool} from '../models';
import {endpoint} from '../models';

export function GetConfig():Promise<config.Config>;

export function GetPotionSubforum():Promise<any>;

export function GetPotionThread():Promise<any>;

export function GetTool():Promise<tool.Tool>;

export function Login(arg1:string,arg2:string,arg3:boolean):Promise<tool.LoginResponse>;

export function SubforumPotionsClub(arg1:Array<string>,arg2:number,arg3:number):Promise<endpoint.SubforumPotionsClubResponse>;

export function UpdatePotionSubforum(arg1:any):Promise<void>;

export function UpdatePotionThread(arg1:any):Promise<void>;
