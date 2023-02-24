/*
Copyright 2023 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in
compliance with the License. You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is
distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing permissions and limitations under the
License.
*/

package internal

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2/dsl/core"
	. "github.com/onsi/gomega"

	"github.com/rh-ecosystem-edge/ztp-pipeline-relocatable/ztp/internal/logging"
)

var _ = Describe("Console", func() {
	var logger logr.Logger

	BeforeEach(func() {
		var err error
		logger, err = logging.NewLogger().
			SetWriter(GinkgoWriter).
			SetLevel(2).
			Build()
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Creation", func() {
		It("Can't be created without a logger", func() {
			console, err := NewConsole().
				SetOut(io.Discard).
				SetErr(io.Discard).
				Build()
			Expect(console).To(BeNil())
			Expect(err).To(HaveOccurred())
			msg := err.Error()
			Expect(msg).To(ContainSubstring("logger"))
			Expect(msg).To(ContainSubstring("mandatory"))
		})

		It("Can't be created without standard output", func() {
			console, err := NewConsole().
				SetLogger(logger).
				SetErr(io.Discard).
				Build()
			Expect(console).To(BeNil())
			Expect(err).To(HaveOccurred())
			msg := err.Error()
			Expect(msg).To(ContainSubstring("output"))
			Expect(msg).To(ContainSubstring("mandatory"))
		})

		It("Can't be created without standard error", func() {
			console, err := NewConsole().
				SetLogger(logger).
				SetOut(io.Discard).
				Build()
			Expect(console).To(BeNil())
			Expect(err).To(HaveOccurred())
			msg := err.Error()
			Expect(msg).To(ContainSubstring("error"))
			Expect(msg).To(ContainSubstring("mandatory"))
		})
	})

	Describe("Usage", func() {
		It("Writes info messages to the output", func() {
			buffer := &bytes.Buffer{}
			multi := io.MultiWriter(buffer, GinkgoWriter)
			console, err := NewConsole().
				SetLogger(logger).
				SetOut(multi).
				SetErr(io.Discard).
				Build()
			Expect(err).ToNot(HaveOccurred())
			console.Info("Hello!")
			Expect(buffer.String()).To(MatchRegexp(`(?m:^I: Hello!\n$)`))
		})

		It("Writes info messages to the log", func() {
			// Create a logger that writes to a buffer, so that we can inspect the
			// messages:
			buffer := &bytes.Buffer{}
			multi := io.MultiWriter(buffer, GinkgoWriter)
			logger, err := logging.NewLogger().
				SetWriter(multi).
				SetLevel(1).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Create the console:
			console, err := NewConsole().
				SetLogger(logger).
				SetOut(io.Discard).
				SetErr(io.Discard).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Verify that it writes the messages to the logs:
			console.Info("Hello!")
			type Msg struct {
				Msg  string `json:"msg"`
				Text string `json:"text"`
			}
			var msg Msg
			err = json.Unmarshal(buffer.Bytes(), &msg)
			Expect(err).ToNot(HaveOccurred())
			Expect(msg.Msg).To(Equal("Console info"))
			Expect(msg.Text).To(Equal("Hello!"))
		})

		It("Writes error messages to the output", func() {
			buffer := &bytes.Buffer{}
			multi := io.MultiWriter(buffer, GinkgoWriter)
			console, err := NewConsole().
				SetLogger(logger).
				SetOut(io.Discard).
				SetErr(multi).
				Build()
			Expect(err).ToNot(HaveOccurred())
			console.Error("Hello!")
			Expect(buffer.String()).To(MatchRegexp(`(?m:^E: Hello!\n$)`))
		})

		It("Writes error messages to the log", func() {
			// Create a logger that writes to a buffer, so that we can inspect the
			// messages:
			buffer := &bytes.Buffer{}
			multi := io.MultiWriter(buffer, GinkgoWriter)
			logger, err := logging.NewLogger().
				SetWriter(multi).
				SetLevel(1).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Create the console:
			console, err := NewConsole().
				SetLogger(logger).
				SetOut(io.Discard).
				SetErr(io.Discard).
				Build()
			Expect(err).ToNot(HaveOccurred())

			// Verify that it writes the messages to the logs:
			console.Error("Hello!")
			type Msg struct {
				Msg  string `json:"msg"`
				Text string `json:"text"`
			}
			var msg Msg
			err = json.Unmarshal(buffer.Bytes(), &msg)
			Expect(err).ToNot(HaveOccurred())
			Expect(msg.Msg).To(Equal("Console error"))
			Expect(msg.Text).To(Equal("Hello!"))
		})
	})
})