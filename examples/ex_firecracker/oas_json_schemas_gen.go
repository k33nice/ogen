// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = math.Mod
	_ = validate.Int{}
)

// WriteJSON implements json.Marshaler.
func (s Balloon) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("amount_mib")
	j.WriteInt(s.AmountMib)
	field.Write("deflate_on_oom")
	j.WriteBool(s.DeflateOnOom)
	if s.StatsPollingIntervalS.Set {
		field.Write("stats_polling_interval_s")
		s.StatsPollingIntervalS.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Balloon json value to io.Writer.
func (s Balloon) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Balloon json value from io.Reader.
func (s *Balloon) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Balloon from json stream.
func (s *Balloon) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "amount_mib":
			s.AmountMib = i.ReadInt()
			return i.Error == nil
		case "deflate_on_oom":
			s.DeflateOnOom = i.ReadBool()
			return i.Error == nil
		case "stats_polling_interval_s":
			s.StatsPollingIntervalS.Reset()
			if err := s.StatsPollingIntervalS.ReadJSON(i); err != nil {
				i.ReportError("Field StatsPollingIntervalS", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s BalloonStats) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("actual_mib")
	j.WriteInt(s.ActualMib)
	field.Write("actual_pages")
	j.WriteInt(s.ActualPages)
	if s.AvailableMemory.Set {
		field.Write("available_memory")
		s.AvailableMemory.WriteJSON(j)
	}
	if s.DiskCaches.Set {
		field.Write("disk_caches")
		s.DiskCaches.WriteJSON(j)
	}
	if s.FreeMemory.Set {
		field.Write("free_memory")
		s.FreeMemory.WriteJSON(j)
	}
	if s.HugetlbAllocations.Set {
		field.Write("hugetlb_allocations")
		s.HugetlbAllocations.WriteJSON(j)
	}
	if s.HugetlbFailures.Set {
		field.Write("hugetlb_failures")
		s.HugetlbFailures.WriteJSON(j)
	}
	if s.MajorFaults.Set {
		field.Write("major_faults")
		s.MajorFaults.WriteJSON(j)
	}
	if s.MinorFaults.Set {
		field.Write("minor_faults")
		s.MinorFaults.WriteJSON(j)
	}
	if s.SwapIn.Set {
		field.Write("swap_in")
		s.SwapIn.WriteJSON(j)
	}
	if s.SwapOut.Set {
		field.Write("swap_out")
		s.SwapOut.WriteJSON(j)
	}
	field.Write("target_mib")
	j.WriteInt(s.TargetMib)
	field.Write("target_pages")
	j.WriteInt(s.TargetPages)
	if s.TotalMemory.Set {
		field.Write("total_memory")
		s.TotalMemory.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes BalloonStats json value to io.Writer.
func (s BalloonStats) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads BalloonStats json value from io.Reader.
func (s *BalloonStats) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads BalloonStats from json stream.
func (s *BalloonStats) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "actual_mib":
			s.ActualMib = i.ReadInt()
			return i.Error == nil
		case "actual_pages":
			s.ActualPages = i.ReadInt()
			return i.Error == nil
		case "available_memory":
			s.AvailableMemory.Reset()
			if err := s.AvailableMemory.ReadJSON(i); err != nil {
				i.ReportError("Field AvailableMemory", err.Error())
				return false
			}
			return true
		case "disk_caches":
			s.DiskCaches.Reset()
			if err := s.DiskCaches.ReadJSON(i); err != nil {
				i.ReportError("Field DiskCaches", err.Error())
				return false
			}
			return true
		case "free_memory":
			s.FreeMemory.Reset()
			if err := s.FreeMemory.ReadJSON(i); err != nil {
				i.ReportError("Field FreeMemory", err.Error())
				return false
			}
			return true
		case "hugetlb_allocations":
			s.HugetlbAllocations.Reset()
			if err := s.HugetlbAllocations.ReadJSON(i); err != nil {
				i.ReportError("Field HugetlbAllocations", err.Error())
				return false
			}
			return true
		case "hugetlb_failures":
			s.HugetlbFailures.Reset()
			if err := s.HugetlbFailures.ReadJSON(i); err != nil {
				i.ReportError("Field HugetlbFailures", err.Error())
				return false
			}
			return true
		case "major_faults":
			s.MajorFaults.Reset()
			if err := s.MajorFaults.ReadJSON(i); err != nil {
				i.ReportError("Field MajorFaults", err.Error())
				return false
			}
			return true
		case "minor_faults":
			s.MinorFaults.Reset()
			if err := s.MinorFaults.ReadJSON(i); err != nil {
				i.ReportError("Field MinorFaults", err.Error())
				return false
			}
			return true
		case "swap_in":
			s.SwapIn.Reset()
			if err := s.SwapIn.ReadJSON(i); err != nil {
				i.ReportError("Field SwapIn", err.Error())
				return false
			}
			return true
		case "swap_out":
			s.SwapOut.Reset()
			if err := s.SwapOut.ReadJSON(i); err != nil {
				i.ReportError("Field SwapOut", err.Error())
				return false
			}
			return true
		case "target_mib":
			s.TargetMib = i.ReadInt()
			return i.Error == nil
		case "target_pages":
			s.TargetPages = i.ReadInt()
			return i.Error == nil
		case "total_memory":
			s.TotalMemory.Reset()
			if err := s.TotalMemory.ReadJSON(i); err != nil {
				i.ReportError("Field TotalMemory", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s BalloonStatsUpdate) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("stats_polling_interval_s")
	j.WriteInt(s.StatsPollingIntervalS)
	j.WriteObjectEnd()
}

// WriteJSONTo writes BalloonStatsUpdate json value to io.Writer.
func (s BalloonStatsUpdate) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads BalloonStatsUpdate json value from io.Reader.
func (s *BalloonStatsUpdate) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads BalloonStatsUpdate from json stream.
func (s *BalloonStatsUpdate) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "stats_polling_interval_s":
			s.StatsPollingIntervalS = i.ReadInt()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s BalloonUpdate) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("amount_mib")
	j.WriteInt(s.AmountMib)
	j.WriteObjectEnd()
}

// WriteJSONTo writes BalloonUpdate json value to io.Writer.
func (s BalloonUpdate) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads BalloonUpdate json value from io.Reader.
func (s *BalloonUpdate) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads BalloonUpdate from json stream.
func (s *BalloonUpdate) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "amount_mib":
			s.AmountMib = i.ReadInt()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s BootSource) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.BootArgs.Set {
		field.Write("boot_args")
		s.BootArgs.WriteJSON(j)
	}
	if s.InitrdPath.Set {
		field.Write("initrd_path")
		s.InitrdPath.WriteJSON(j)
	}
	field.Write("kernel_image_path")
	j.WriteString(s.KernelImagePath)
	j.WriteObjectEnd()
}

// WriteJSONTo writes BootSource json value to io.Writer.
func (s BootSource) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads BootSource json value from io.Reader.
func (s *BootSource) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads BootSource from json stream.
func (s *BootSource) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "boot_args":
			s.BootArgs.Reset()
			if err := s.BootArgs.ReadJSON(i); err != nil {
				i.ReportError("Field BootArgs", err.Error())
				return false
			}
			return true
		case "initrd_path":
			s.InitrdPath.Reset()
			if err := s.InitrdPath.ReadJSON(i); err != nil {
				i.ReportError("Field InitrdPath", err.Error())
				return false
			}
			return true
		case "kernel_image_path":
			s.KernelImagePath = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Drive) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.CacheType.Set {
		field.Write("cache_type")
		s.CacheType.WriteJSON(j)
	}
	field.Write("drive_id")
	j.WriteString(s.DriveID)
	field.Write("is_read_only")
	j.WriteBool(s.IsReadOnly)
	field.Write("is_root_device")
	j.WriteBool(s.IsRootDevice)
	if s.Partuuid.Set {
		field.Write("partuuid")
		s.Partuuid.WriteJSON(j)
	}
	field.Write("path_on_host")
	j.WriteString(s.PathOnHost)
	// Unsupported kind "pointer" for field "RateLimiter".
	j.WriteObjectEnd()
}

// WriteJSONTo writes Drive json value to io.Writer.
func (s Drive) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Drive json value from io.Reader.
func (s *Drive) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Drive from json stream.
func (s *Drive) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "cache_type":
			s.CacheType.Reset()
			if err := s.CacheType.ReadJSON(i); err != nil {
				i.ReportError("Field CacheType", err.Error())
				return false
			}
			return true
		case "drive_id":
			s.DriveID = i.ReadString()
			return i.Error == nil
		case "is_read_only":
			s.IsReadOnly = i.ReadBool()
			return i.Error == nil
		case "is_root_device":
			s.IsRootDevice = i.ReadBool()
			return i.Error == nil
		case "partuuid":
			s.Partuuid.Reset()
			if err := s.Partuuid.ReadJSON(i); err != nil {
				i.ReportError("Field Partuuid", err.Error())
				return false
			}
			return true
		case "path_on_host":
			s.PathOnHost = i.ReadString()
			return i.Error == nil
		case "rate_limiter":
			// Unsupported kind "pointer" for field "RateLimiter".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Error) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.FaultMessage.Set {
		field.Write("fault_message")
		s.FaultMessage.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Error json value to io.Writer.
func (s Error) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Error json value from io.Reader.
func (s *Error) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Error from json stream.
func (s *Error) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "fault_message":
			s.FaultMessage.Reset()
			if err := s.FaultMessage.ReadJSON(i); err != nil {
				i.ReportError("Field FaultMessage", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s ErrorStatusCode) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	j.WriteObjectEnd()
}

// WriteJSONTo writes ErrorStatusCode json value to io.Writer.
func (s ErrorStatusCode) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads ErrorStatusCode json value from io.Reader.
func (s *ErrorStatusCode) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads ErrorStatusCode from json stream.
func (s *ErrorStatusCode) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s FullVmConfiguration) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	// Unsupported kind "pointer" for field "BalloonDevice".
	// Unsupported kind "pointer" for field "BlockDevices".
	// Unsupported kind "pointer" for field "BootSource".
	// Unsupported kind "pointer" for field "Logger".
	// Unsupported kind "pointer" for field "MachineConfig".
	// Unsupported kind "pointer" for field "Metrics".
	// Unsupported kind "pointer" for field "MmdsConfig".
	// Unsupported kind "pointer" for field "NetDevices".
	// Unsupported kind "pointer" for field "VsockDevice".
	j.WriteObjectEnd()
}

// WriteJSONTo writes FullVmConfiguration json value to io.Writer.
func (s FullVmConfiguration) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads FullVmConfiguration json value from io.Reader.
func (s *FullVmConfiguration) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads FullVmConfiguration from json stream.
func (s *FullVmConfiguration) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "balloon_device":
			// Unsupported kind "pointer" for field "BalloonDevice".
			i.Skip()
			return true
		case "block_devices":
			// Unsupported kind "pointer" for field "BlockDevices".
			i.Skip()
			return true
		case "boot_source":
			// Unsupported kind "pointer" for field "BootSource".
			i.Skip()
			return true
		case "logger":
			// Unsupported kind "pointer" for field "Logger".
			i.Skip()
			return true
		case "machine_config":
			// Unsupported kind "pointer" for field "MachineConfig".
			i.Skip()
			return true
		case "metrics":
			// Unsupported kind "pointer" for field "Metrics".
			i.Skip()
			return true
		case "mmds_config":
			// Unsupported kind "pointer" for field "MmdsConfig".
			i.Skip()
			return true
		case "net_devices":
			// Unsupported kind "pointer" for field "NetDevices".
			i.Skip()
			return true
		case "vsock_device":
			// Unsupported kind "pointer" for field "VsockDevice".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s InstanceActionInfo) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("action_type")
	j.WriteString(s.ActionType)
	j.WriteObjectEnd()
}

// WriteJSONTo writes InstanceActionInfo json value to io.Writer.
func (s InstanceActionInfo) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads InstanceActionInfo json value from io.Reader.
func (s *InstanceActionInfo) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads InstanceActionInfo from json stream.
func (s *InstanceActionInfo) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "action_type":
			s.ActionType = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s InstanceInfo) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("app_name")
	j.WriteString(s.AppName)
	field.Write("id")
	j.WriteString(s.ID)
	field.Write("state")
	j.WriteString(s.State)
	field.Write("vmm_version")
	j.WriteString(s.VmmVersion)
	j.WriteObjectEnd()
}

// WriteJSONTo writes InstanceInfo json value to io.Writer.
func (s InstanceInfo) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads InstanceInfo json value from io.Reader.
func (s *InstanceInfo) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads InstanceInfo from json stream.
func (s *InstanceInfo) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "app_name":
			s.AppName = i.ReadString()
			return i.Error == nil
		case "id":
			s.ID = i.ReadString()
			return i.Error == nil
		case "state":
			s.State = i.ReadString()
			return i.Error == nil
		case "vmm_version":
			s.VmmVersion = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Logger) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	// Unsupported kind "pointer" for field "Level".
	field.Write("log_path")
	j.WriteString(s.LogPath)
	if s.ShowLevel.Set {
		field.Write("show_level")
		s.ShowLevel.WriteJSON(j)
	}
	if s.ShowLogOrigin.Set {
		field.Write("show_log_origin")
		s.ShowLogOrigin.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes Logger json value to io.Writer.
func (s Logger) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Logger json value from io.Reader.
func (s *Logger) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Logger from json stream.
func (s *Logger) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "level":
			// Unsupported kind "pointer" for field "Level".
			i.Skip()
			return true
		case "log_path":
			s.LogPath = i.ReadString()
			return i.Error == nil
		case "show_level":
			s.ShowLevel.Reset()
			if err := s.ShowLevel.ReadJSON(i); err != nil {
				i.ReportError("Field ShowLevel", err.Error())
				return false
			}
			return true
		case "show_log_origin":
			s.ShowLogOrigin.Reset()
			if err := s.ShowLogOrigin.ReadJSON(i); err != nil {
				i.ReportError("Field ShowLogOrigin", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s MachineConfiguration) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	// Unsupported kind "pointer" for field "CPUTemplate".
	field.Write("ht_enabled")
	j.WriteBool(s.HtEnabled)
	field.Write("mem_size_mib")
	j.WriteInt(s.MemSizeMib)
	if s.TrackDirtyPages.Set {
		field.Write("track_dirty_pages")
		s.TrackDirtyPages.WriteJSON(j)
	}
	field.Write("vcpu_count")
	j.WriteInt(s.VcpuCount)
	j.WriteObjectEnd()
}

// WriteJSONTo writes MachineConfiguration json value to io.Writer.
func (s MachineConfiguration) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads MachineConfiguration json value from io.Reader.
func (s *MachineConfiguration) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads MachineConfiguration from json stream.
func (s *MachineConfiguration) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "cpu_template":
			// Unsupported kind "pointer" for field "CPUTemplate".
			i.Skip()
			return true
		case "ht_enabled":
			s.HtEnabled = i.ReadBool()
			return i.Error == nil
		case "mem_size_mib":
			s.MemSizeMib = i.ReadInt()
			return i.Error == nil
		case "track_dirty_pages":
			s.TrackDirtyPages.Reset()
			if err := s.TrackDirtyPages.ReadJSON(i); err != nil {
				i.ReportError("Field TrackDirtyPages", err.Error())
				return false
			}
			return true
		case "vcpu_count":
			s.VcpuCount = i.ReadInt()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Metrics) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("metrics_path")
	j.WriteString(s.MetricsPath)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Metrics json value to io.Writer.
func (s Metrics) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Metrics json value from io.Reader.
func (s *Metrics) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Metrics from json stream.
func (s *Metrics) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "metrics_path":
			s.MetricsPath = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s MmdsConfig) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.Ipv4Address.Set {
		field.Write("ipv4_address")
		s.Ipv4Address.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes MmdsConfig json value to io.Writer.
func (s MmdsConfig) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads MmdsConfig json value from io.Reader.
func (s *MmdsConfig) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads MmdsConfig from json stream.
func (s *MmdsConfig) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "ipv4_address":
			s.Ipv4Address.Reset()
			if err := s.Ipv4Address.ReadJSON(i); err != nil {
				i.ReportError("Field Ipv4Address", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s NetworkInterface) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.AllowMmdsRequests.Set {
		field.Write("allow_mmds_requests")
		s.AllowMmdsRequests.WriteJSON(j)
	}
	if s.GuestMAC.Set {
		field.Write("guest_mac")
		s.GuestMAC.WriteJSON(j)
	}
	field.Write("host_dev_name")
	j.WriteString(s.HostDevName)
	field.Write("iface_id")
	j.WriteString(s.IfaceID)
	// Unsupported kind "pointer" for field "RxRateLimiter".
	// Unsupported kind "pointer" for field "TxRateLimiter".
	j.WriteObjectEnd()
}

// WriteJSONTo writes NetworkInterface json value to io.Writer.
func (s NetworkInterface) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads NetworkInterface json value from io.Reader.
func (s *NetworkInterface) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads NetworkInterface from json stream.
func (s *NetworkInterface) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "allow_mmds_requests":
			s.AllowMmdsRequests.Reset()
			if err := s.AllowMmdsRequests.ReadJSON(i); err != nil {
				i.ReportError("Field AllowMmdsRequests", err.Error())
				return false
			}
			return true
		case "guest_mac":
			s.GuestMAC.Reset()
			if err := s.GuestMAC.ReadJSON(i); err != nil {
				i.ReportError("Field GuestMAC", err.Error())
				return false
			}
			return true
		case "host_dev_name":
			s.HostDevName = i.ReadString()
			return i.Error == nil
		case "iface_id":
			s.IfaceID = i.ReadString()
			return i.Error == nil
		case "rx_rate_limiter":
			// Unsupported kind "pointer" for field "RxRateLimiter".
			i.Skip()
			return true
		case "tx_rate_limiter":
			// Unsupported kind "pointer" for field "TxRateLimiter".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PartialDrive) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("drive_id")
	j.WriteString(s.DriveID)
	if s.PathOnHost.Set {
		field.Write("path_on_host")
		s.PathOnHost.WriteJSON(j)
	}
	// Unsupported kind "pointer" for field "RateLimiter".
	j.WriteObjectEnd()
}

// WriteJSONTo writes PartialDrive json value to io.Writer.
func (s PartialDrive) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PartialDrive json value from io.Reader.
func (s *PartialDrive) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PartialDrive from json stream.
func (s *PartialDrive) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "drive_id":
			s.DriveID = i.ReadString()
			return i.Error == nil
		case "path_on_host":
			s.PathOnHost.Reset()
			if err := s.PathOnHost.ReadJSON(i); err != nil {
				i.ReportError("Field PathOnHost", err.Error())
				return false
			}
			return true
		case "rate_limiter":
			// Unsupported kind "pointer" for field "RateLimiter".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s PartialNetworkInterface) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("iface_id")
	j.WriteString(s.IfaceID)
	// Unsupported kind "pointer" for field "RxRateLimiter".
	// Unsupported kind "pointer" for field "TxRateLimiter".
	j.WriteObjectEnd()
}

// WriteJSONTo writes PartialNetworkInterface json value to io.Writer.
func (s PartialNetworkInterface) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads PartialNetworkInterface json value from io.Reader.
func (s *PartialNetworkInterface) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads PartialNetworkInterface from json stream.
func (s *PartialNetworkInterface) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "iface_id":
			s.IfaceID = i.ReadString()
			return i.Error == nil
		case "rx_rate_limiter":
			// Unsupported kind "pointer" for field "RxRateLimiter".
			i.Skip()
			return true
		case "tx_rate_limiter":
			// Unsupported kind "pointer" for field "TxRateLimiter".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s RateLimiter) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	// Unsupported kind "pointer" for field "Bandwidth".
	// Unsupported kind "pointer" for field "Ops".
	j.WriteObjectEnd()
}

// WriteJSONTo writes RateLimiter json value to io.Writer.
func (s RateLimiter) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads RateLimiter json value from io.Reader.
func (s *RateLimiter) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads RateLimiter from json stream.
func (s *RateLimiter) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "bandwidth":
			// Unsupported kind "pointer" for field "Bandwidth".
			i.Skip()
			return true
		case "ops":
			// Unsupported kind "pointer" for field "Ops".
			i.Skip()
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s SnapshotCreateParams) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("mem_file_path")
	j.WriteString(s.MemFilePath)
	field.Write("snapshot_path")
	j.WriteString(s.SnapshotPath)
	// Unsupported kind "pointer" for field "SnapshotType".
	if s.Version.Set {
		field.Write("version")
		s.Version.WriteJSON(j)
	}
	j.WriteObjectEnd()
}

// WriteJSONTo writes SnapshotCreateParams json value to io.Writer.
func (s SnapshotCreateParams) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads SnapshotCreateParams json value from io.Reader.
func (s *SnapshotCreateParams) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads SnapshotCreateParams from json stream.
func (s *SnapshotCreateParams) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "mem_file_path":
			s.MemFilePath = i.ReadString()
			return i.Error == nil
		case "snapshot_path":
			s.SnapshotPath = i.ReadString()
			return i.Error == nil
		case "snapshot_type":
			// Unsupported kind "pointer" for field "SnapshotType".
			i.Skip()
			return true
		case "version":
			s.Version.Reset()
			if err := s.Version.ReadJSON(i); err != nil {
				i.ReportError("Field Version", err.Error())
				return false
			}
			return true
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s SnapshotLoadParams) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.EnableDiffSnapshots.Set {
		field.Write("enable_diff_snapshots")
		s.EnableDiffSnapshots.WriteJSON(j)
	}
	field.Write("mem_file_path")
	j.WriteString(s.MemFilePath)
	if s.ResumeVM.Set {
		field.Write("resume_vm")
		s.ResumeVM.WriteJSON(j)
	}
	field.Write("snapshot_path")
	j.WriteString(s.SnapshotPath)
	j.WriteObjectEnd()
}

// WriteJSONTo writes SnapshotLoadParams json value to io.Writer.
func (s SnapshotLoadParams) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads SnapshotLoadParams json value from io.Reader.
func (s *SnapshotLoadParams) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads SnapshotLoadParams from json stream.
func (s *SnapshotLoadParams) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "enable_diff_snapshots":
			s.EnableDiffSnapshots.Reset()
			if err := s.EnableDiffSnapshots.ReadJSON(i); err != nil {
				i.ReportError("Field EnableDiffSnapshots", err.Error())
				return false
			}
			return true
		case "mem_file_path":
			s.MemFilePath = i.ReadString()
			return i.Error == nil
		case "resume_vm":
			s.ResumeVM.Reset()
			if err := s.ResumeVM.ReadJSON(i); err != nil {
				i.ReportError("Field ResumeVM", err.Error())
				return false
			}
			return true
		case "snapshot_path":
			s.SnapshotPath = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s TokenBucket) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	if s.OneTimeBurst.Set {
		field.Write("one_time_burst")
		s.OneTimeBurst.WriteJSON(j)
	}
	field.Write("refill_time")
	j.WriteInt64(s.RefillTime)
	field.Write("size")
	j.WriteInt64(s.Size)
	j.WriteObjectEnd()
}

// WriteJSONTo writes TokenBucket json value to io.Writer.
func (s TokenBucket) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads TokenBucket json value from io.Reader.
func (s *TokenBucket) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads TokenBucket from json stream.
func (s *TokenBucket) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "one_time_burst":
			s.OneTimeBurst.Reset()
			if err := s.OneTimeBurst.ReadJSON(i); err != nil {
				i.ReportError("Field OneTimeBurst", err.Error())
				return false
			}
			return true
		case "refill_time":
			s.RefillTime = i.ReadInt64()
			return i.Error == nil
		case "size":
			s.Size = i.ReadInt64()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s VM) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("state")
	j.WriteString(s.State)
	j.WriteObjectEnd()
}

// WriteJSONTo writes VM json value to io.Writer.
func (s VM) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads VM json value from io.Reader.
func (s *VM) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads VM from json stream.
func (s *VM) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "state":
			s.State = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}

// WriteJSON implements json.Marshaler.
func (s Vsock) WriteJSON(j *json.Stream) {
	j.WriteObjectStart()
	field := json.NewFieldWriter(j)
	defer field.Reset()
	field.Write("guest_cid")
	j.WriteInt(s.GuestCid)
	field.Write("uds_path")
	j.WriteString(s.UdsPath)
	field.Write("vsock_id")
	j.WriteString(s.VsockID)
	j.WriteObjectEnd()
}

// WriteJSONTo writes Vsock json value to io.Writer.
func (s Vsock) WriteJSONTo(w io.Writer) error {
	j := json.NewStream(w)
	s.WriteJSON(j)
	return j.Flush()
}

// ReadJSONFrom reads Vsock json value from io.Reader.
func (s *Vsock) ReadJSONFrom(r io.Reader) error {
	buf := json.GetBuffer()
	defer json.PutBuffer(buf)

	if _, err := buf.ReadFrom(r); err != nil {
		return err
	}
	i := json.GetIterator()
	i.ResetBytes(buf.Bytes())
	defer json.PutIterator(i)

	return s.ReadJSON(i)
}

// ReadJSON reads Vsock from json stream.
func (s *Vsock) ReadJSON(i *json.Iterator) error {
	i.ReadObjectCB(func(i *json.Iterator, k string) bool {
		switch k {
		case "guest_cid":
			s.GuestCid = i.ReadInt()
			return i.Error == nil
		case "uds_path":
			s.UdsPath = i.ReadString()
			return i.Error == nil
		case "vsock_id":
			s.VsockID = i.ReadString()
			return i.Error == nil
		default:
			i.Skip()
			return true
		}
	})
	return i.Error
}
