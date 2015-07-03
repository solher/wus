package main

import (
	"runtime"

	_ "github.com/solher/wus/domain"
	_ "github.com/solher/wus/ressources"
	"github.com/solher/zest"
	"github.com/solher/zest/domain"
	"github.com/solher/zest/ressources"
	"github.com/solher/zest/usecases"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	App := zest.Classic()
	App.AfterBuild = afterBuild
	App.UserSeedDatabase = userSeedDatabase

	App.Run()
}

func afterBuild(z *zest.Zest) error {
	type dependencies struct {
		RouteDir *usecases.RouteDirectory
	}

	d := &dependencies{}
	err := z.Injector.Get(d)
	if err != nil {
		return err
	}

	routes := d.RouteDir.Routes

	for k, v := range routes {
		if k.Ressource == "urls" && k.Method == "RedirectUrl" {
			continue
		}

		v.Path = "/api" + v.Path
		routes[k] = v
	}

	return nil
}

func userSeedDatabase(z *zest.Zest) error {
	type dependencies struct {
		AccountGuestInter *ressources.AccountGuestInter
		PermissionInter   *usecases.PermissionInter
	}

	d := &dependencies{}
	err := z.Injector.Get(d)
	if err != nil {
		return err
	}

	user := &domain.User{
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "admin",
		Email:     "admin@wid.la",
	}

	account, err := d.AccountGuestInter.Signup(user)
	if err != nil {
		return err
	}

	err = d.PermissionInter.SetRole(account.ID, "Admin")
	if err != nil {
		return err
	}

	err = d.PermissionInter.SetAcl("urls", "RedirectUrl", "Anyone")
	if err != nil {
		return err
	}

	return nil
}
