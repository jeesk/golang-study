package error_handle

import (
	"errors"
	"fmt"
	"testing"
)

type Demo struct {
}

func TestError1(t *testing.T) {
	// 创建一个error 对象
	err := errors.New("我是错误")
	errx := errors.New("我是错误")
	// warper
	_ = fmt.Errorf("包裹了错误%w", nil)
	/*	err2 := errors.Unwrap(errW)*/
	if errors.Is(err, errx) {
		fmt.Print("相等")
	}

}

/*func TestText(t *testing.T) {
	books := []string{
		"Book1",
		"Book222222",
		"Book3333333333",
	}

	for _, bookName := range books {
		err := searchBook(bookName)

		// 特殊业务场景：如果发现书被借走了，下次再来就行了，不需要作为错误处理
		if err != nil {
			// 提取error这个interface底层的错误码，一般在API的返回前才提取
			// As - 获取错误的具体实现
			var myError = new(CustomError)
			// Is - 判断错误是否为指定类型
			if errors.Cause(err) == NewCustomError(BookHasBeenBorrowedError) {
				fmt.Printf("类型相同")
			}
			if errors.Is(err, NewCustomError(BookHasBeenBorrowedError)) {
				fmt.Printf("IS中的信息：%s 已经被借走了, 只需按Info处理!\n", bookName)
				err = nil
			} else {
				// 如果已有堆栈信息，应调用WithMessage方法
				//newErr := errors.WithMessage(err, "WithMessage err")
				fmt.Printf("IS中的信息：%s 未找到，应该按Error处理! ,newErr is %s\n", bookName, myError)
			}
			// As - 解析错误内容
			if errors.As(err, &myError) {
				fmt.Printf("AS中的信息：当前书为: %s ,error code is %d, message is %s\n", bookName, myError.Code, myError.TagMessage)
			}

			// 特殊场景，指定错误(ErrorBookHasBeenBorrowed)时，打印即可，不返回错误

		}
	}
}

func searchBook(bookName string) error {
	// 1 发现图书馆不存在这本书 - 认为是错误，需要打印详细的错误信息
	if len(bookName) > 10 {
		return NewCustomError(BookHasBeenBorrowedError)
	} else if len(bookName) > 6 {
		// 2 发现书被借走了 - 打印一下被接走的提示即可，不认为是错误
		return NewCustomError(BookHasBeenBorrowedError)
	}
	// 3 找到书 - 不需要任何处理
	return nil
}
*/
