package main

import (
		"fmt"
		"strings"
		"time"
		
		"golang.org/x/sys/windows/registry" // 레지스트리 패키지 추가

		"github.com/shirou/gopsutil/process"
)

func main() {
		fmt.Println("dddd")
		fmt.Println("5초 뒤에 우회가 시작됩니다.")
		// 일시 중지할 프로세스 이름 리스트
		// 5초 동안 멈춥니다.
		duration := 5 * time.Second
		time.Sleep(duration)
		processNames := []string{"iAgent.exe", "iAgent32.exe", "iWatcher.exe", "iService.exe", "systemama.exe", "systemams.exe"}

		// 현재 실행 중인 모든 프로세스 목록 가져오기
		processList, err := process.Processes()
		if err != nil{
			fmt.Printf("프로세스 목록을 가져오는 동안 오류 발생: %v\n", err)
			return
		}

		//입력한 각 프로세스 이름에 대해 프로세스를 찾아서 중지
		for _, processName := range processNames {
			for _, p := range processList {
				name, err := p.Name()
				if err != nil {
					continue
				}

				if strings.Contains(strings.ToLower(name), strings.ToLower(processName)) {
					err := p.Suspend()
					if err == nil {
						fmt.printf("프로세스 %s (PID %d) 우회 성공.\n", name, p.Pid)
					} else {
						fmt.Printf("프로세스 %s (PID %d) 우회 실패 관리자 권한으로 실행해주세요: %v\n", name, p.Pid, err)
					}
				}
			}
		}

		// 프록시 설정을 비활성화 할 레지스트리 키 경로
		keyPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`

		// 레지스트리 키 열기 (키 경로에 대한 쓰기 권한 필요)
		key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.SET_VALUE)
		if err != nil {
			fmt.Printf("레지스트리 키 열기 실패: %v\n", err)
			return
		}
		defer key.Close()

		// 프록시 설정 비활성화
		err = key.SetDWordValue("ProxyEnable", 0)
		if err != nil {
			fmt.Printf("프록시 설정 비활성화 실패: %v\n", err)
			return
		}

		fmt.Println("ㅍ프록시 설정이 비활성화 되었습니다.")
		fmt.Println("프로그램을 종료하려면 언터 키를 누르세요.")
		fmt.Scanln() // 엔터 키를 누를 떄까지 대기
	}