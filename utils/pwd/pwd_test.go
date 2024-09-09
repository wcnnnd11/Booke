package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$MowiThF2eBlVMNCPkTwAzuqbcOQb.WvSybLGsUn9tLG9TBHQAYxUm", "1234"))
}
