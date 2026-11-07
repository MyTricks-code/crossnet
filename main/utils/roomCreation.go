package utils

import (
	"errors"
	"fmt"
	"net"
)

type Room struct{
	owner net.Conn
	Peers []net.Conn
}

var rooms =make(map[int]*Room)

func CreateRoom(id int, peer net.Conn) error{	
	_, exists:=rooms[id]
	if exists==true{
		fmt.Println("Room already exists")
		return errors.New("Room already exists")
	}else{
		rooms[id]=&Room{owner: peer}
		fmt.Println("Room already exists")
		return nil;
	}
}

func DeleteRoom(id int, peer net.Conn) error{
	_, exists:=rooms[id]
	if exists==true{
		if rooms[id].owner==peer{
			for _, peer :=range rooms[id].Peers{
				peer.Close()
			}
			delete(rooms, id)
			fmt.Println("Room deleted")
			return nil;
		}else{
			fmt.Println("Ownership mismatched")
			return errors.New("Ownership mismatched")
		}
	}else{
		fmt.Println("No room found")
		return errors.New("No room found")
	}
}

func AddPeer(id int, peer net.Conn) error{
	_,exists:=rooms[id];
	if exists!=true {
		fmt.Println("No room found")
		return errors.New("No room found")
	}
	rooms[id].Peers = append(rooms[id].Peers, peer)
	return nil;
}

// func DeletePeer(id int, peer net.Conn) error{
// 	_, exists:=rooms[id];
// 	if exists!=true {
// 		fmt.Println("No room found")
// 		return errors.New("No room found")
// 	}
// 	peerList:=rooms[id].Peers
// 	delete(peerList, peer)
// 	return nil;
// }