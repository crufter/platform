import { Injectable } from "@angular/core";
import {
  CanActivate,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
  UrlTree
} from "@angular/router";
import { Observable } from "rxjs";
import { UserService } from "./user.service";
import { environment } from "../environments/environment";

@Injectable({
  providedIn: "root"
})
export class AuthGuard implements CanActivate {
  constructor(private us: UserService) {}

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ):
    | Observable<boolean | UrlTree>
    | Promise<boolean | UrlTree>
    | boolean
    | UrlTree {
    if (this.us.loggedIn()) {
      return true
    }
    console.log("subscribin")
    return new Observable<boolean>((observer) => {
      this.us.isUserLoggedIn.subscribe(loggedIn => {
        if (loggedIn) {
          
          observer.next(true)
        } else {
          console.log("auth guard")
          confirm("redirect") ? window.location.href = environment.backendUrl + "/v1/github/login" : console.log("stopping")
          observer.next(false)
        }
        observer.complete()
      })
    });
  }
}
